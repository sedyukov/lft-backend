package blockchain

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
	contracts "github.com/sedyukov/lft-backend/contracts/interfaces"
	lftcontrollers "github.com/sedyukov/lft-backend/internal/controllers/lft"
	lftdb "github.com/sedyukov/lft-backend/internal/database/lft"
	"golang.org/x/sync/errgroup"
)

const (
	TxReceiptsBatchSize = 16
	RequestTimeout      = 32 * time.Second
	RequestRetryDelay   = 32 * time.Millisecond
	historyBlockBatch   = 1000
)

// Monitor interface
type Monitor interface {
	Start(ctx context.Context, client *ethclient.Client, logger zerolog.Logger) error
	StartRpc(ctx context.Context, client *ethclient.Client, logger zerolog.Logger) error
}

type monitor struct {
	contractAddress string
	client          ethclient.Client
	logger          zerolog.Logger
}

// AllowanceChangedEvent struct
// type RewardReferralEvent struct {
// 	Trader      string   `json:"trader"`
// 	Refferal    string   `json:"refferal"`
// 	Level       uint8    `json:"level"`
// 	Amount      *big.Int `json:"amount"`
// 	BlockNumber uint64   `json:"block_number"`
// }

// OwnershipTransferredEvent struct
type OwnershipTransferredEvent struct {
	Event         string    `json:"event_type"`
	PreviousOwner string    `json:"previous_owner"`
	NewOwner      string    `json:"new_owner"`
	BlockNumber   uint64    `json:"block_number"`
	Timestamp     time.Time `json:"timestamp"`
}

// NewMonitor returns a new runner instance
func NewMonitor(contractAddress string, logger zerolog.Logger) Monitor {
	return &monitor{
		contractAddress: contractAddress,
	}
}

// Start register to listen blockchain events
func (m *monitor) Start(ctx context.Context, client *ethclient.Client, logger zerolog.Logger) error {
	m.client = *client
	m.logger = logger
	logger.Info().Msgf("Start monitoring at %s", m.contractAddress)

	err := validateContractAddress(ctx, client, m.contractAddress)
	if err != nil {
		logger.Error().Msg("Contract address validation failed")
		return err
	}

	contract, err := contracts.NewContract(common.HexToAddress(m.contractAddress), client)
	if err != nil {
		logger.Error().Msg("Contract creation failed")
		return err
	}
	logger.Info().Msg("Contract instance created successfully")
	eg := new(errgroup.Group)
	eg.Go(func() error {
		return m.watchRewardReferralEvent(ctx, contract)
	})
	eg.Go(func() error {
		return m.watchOwnershipTransferred(ctx, contract)
	})
	if err = eg.Wait(); err != nil {
		logger.Error().Msg("Events error group throw error")
		return err
	}
	return nil
}

func (m *monitor) StartRpc(ctx context.Context, client *ethclient.Client, logger zerolog.Logger) error {
	logger.Info().Msgf("Start monitoring at %s", m.contractAddress)

	m.client = *client
	m.logger = logger

	err := validateContractAddress(ctx, client, m.contractAddress)
	if err != nil {
		logger.Error().Msg("Contract address validation failed")
		return err
	}

	hexedAddress := common.HexToAddress(m.contractAddress)

	contract, err := contracts.NewContract(hexedAddress, client)
	if err != nil {
		logger.Error().Msg("Contract creation failed")
		return err
	}
	logger.Info().Msg("Contract instance created successfully")

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		logger.Error().Msg("Failed when retrieving last block")
		return err
	}
	currentBlock := header.Number.Uint64()

	lastBlockFromDb, err := strconv.ParseInt(lftdb.GetLastBlock(), 0, 64)
	if err != nil {
		logger.Error().Msg("Failed when trying to convert last parsed block to int")
		return err
	}
	blockStart := uint64(lastBlockFromDb)

	if blockStart < currentBlock {
		for i := blockStart; i < currentBlock; i += historyBlockBatch {
			currentBlockEnd := i + historyBlockBatch - 1

			if currentBlockEnd >= currentBlock {
				currentBlockEnd = currentBlock - 1
			}

			query := &bind.FilterOpts{
				Start: i,
				End:   &currentBlockEnd,
			}

			eventsIterator, err := contract.FilterRewardReferral(query, nil, nil, nil)
			if err != nil {
				logger.Error().Msg("Get FilterRewardReferral failed")
				return err
			}

			m.logger.Info().Msgf(
				fmt.Sprintf("Fetched batch from %d to %d", i, currentBlockEnd),
			)

			m.parseRewardReferralEvent(eventsIterator)
			lftdb.UpdateLastBlock(strconv.FormatUint(currentBlockEnd, 10))
		}
	}

	for i := currentBlock; true; i++ {
		block := m.fetchBlock(int64(i))

		if block == nil {
			return nil
		}

		endBlock := currentBlock + 1
		logger.Info().Uint64("endBlock", endBlock)
		query := &bind.FilterOpts{
			Start: i,
			End:   &i,
		}

		eventsIterator, err := contract.FilterRewardReferral(query, nil, nil, nil)
		if err != nil {
			logger.Error().Msg("Get FilterRewardReferral failed")
			return err
		}

		m.parseRewardReferralEvent(eventsIterator)
		lftdb.UpdateLastBlock(strconv.FormatUint(i, 10))
	}

	return nil
}

func (m *monitor) fetchBlock(height int64) *types.Header {
	// Request until get block
	for first, start, deadline := true, time.Now(), time.Now().Add(RequestTimeout); true; first = false {
		// Request block
		result, err := m.client.HeaderByNumber(context.Background(), new(big.Int).SetInt64(height))
		if err == nil {
			if !first {
				m.logger.Info().Msgf(
					fmt.Sprintf("Fetched block (after %s), height %d", DurationToString(time.Since(start)), height),
				)
			} else {
				m.logger.Info().Msgf(
					fmt.Sprintf("Fetched block (%s), height %d", DurationToString(time.Since(start)), height),
				)
			}
			return result
		}
		// Stop trying when the deadline is reached
		if time.Now().After(deadline) {
			return nil
		}
		// Sleep some time before next try
		time.Sleep(RequestRetryDelay)
	}

	return nil
}

func (m *monitor) watchRewardReferralEvent(ctx context.Context, contract *contracts.Contract) error {
	events := make(chan *contracts.ContractRewardReferral)
	opts := &bind.WatchOpts{
		Start:   nil,
		Context: ctx,
	}
	subscription, err := contract.WatchRewardReferral(opts, events, nil, nil, nil)
	if err != nil {
		return err
	}
	defer subscription.Unsubscribe()

	for {
		select {
		case <-ctx.Done():
			return nil
		case errChan := <-subscription.Err():
			return errChan
		case event := <-events:
			j, _ := json.MarshalIndent(
				lftcontrollers.RewardReferralEvent{
					Trader:      event.Trader.Hex(),
					Refferal:    event.Referral.Hex(),
					Level:       event.Level,
					Amount:      event.Amount,
					BlockNumber: event.Raw.BlockNumber,
				},
				"",
				"  ",
			)
			fmt.Println(string(j))
		}
	}
}

func (m *monitor) watchOwnershipTransferred(ctx context.Context, contract *contracts.Contract) error {
	events := make(chan *contracts.ContractOwnershipTransferred)
	opts := &bind.WatchOpts{
		Start:   nil,
		Context: ctx,
	}
	subscription, err := contract.WatchOwnershipTransferred(opts, events, nil, nil)
	if err != nil {
		return err
	}
	defer subscription.Unsubscribe()

	for {
		select {
		case <-ctx.Done():
			return nil
		case errChan := <-subscription.Err():
			return errChan
		case event := <-events:
			j, _ := json.MarshalIndent(
				OwnershipTransferredEvent{
					Event:         "OwnershipTransferred",
					PreviousOwner: event.PreviousOwner.Hex(),
					NewOwner:      event.NewOwner.Hex(),
					BlockNumber:   event.Raw.BlockNumber,
					Timestamp:     time.Now(),
				},
				"",
				"  ",
			)
			fmt.Println(string(j))
		}
	}
}

func (m *monitor) parseRewardReferralEvent(eventsIterator *contracts.ContractRewardReferralIterator) error {
	for eventsIterator.Next() {
		event := eventsIterator.Event
		rre := lftcontrollers.RewardReferralEvent{
			Trader:      event.Trader.Hex(),
			Refferal:    event.Referral.Hex(),
			Level:       event.Level,
			Amount:      event.Amount,
			BlockNumber: event.Raw.BlockNumber,
		}
		lftcontrollers.CreateRewardRefferal(rre)
	}

	if eventsIterator.Error() != nil {
		return eventsIterator.Error()
	}

	return nil
}
