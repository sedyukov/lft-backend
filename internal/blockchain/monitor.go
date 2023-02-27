package blockchain

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
	contracts "github.com/sedyukov/lft-backend/contracts/interfaces"
	"golang.org/x/sync/errgroup"
)

// Monitor interface
type Monitor interface {
	Start(ctx context.Context, client *ethclient.Client, logger zerolog.Logger) error
	StartRpc(ctx context.Context, client *ethclient.Client, logger zerolog.Logger) error
}

type monitor struct {
	contractAddress string
}

// AllowanceChangedEvent struct
type RewardReferralEvent struct {
	Trader      string   `json:"trader"`
	Refferal    string   `json:"refferal"`
	Level       uint8    `json:"level"`
	Amount      *big.Int `json:"amount"`
	BlockNumber uint64   `json:"block_number"`
}

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

	// TODO: implement logic for block start
	blockStart := uint64(27262097)
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		logger.Error().Msg("Failed when retrieving last block")
		return err
	}

	currentBlock := header.Number.Uint64()

	if blockStart < currentBlock {
		// query := ethereum.FilterQuery{
		// 	FromBlock: big.NewInt(blockStart),
		// 	ToBlock:   big.NewInt(2394201),
		// 	Addresses: []common.Address{
		// 		hexedAddress,
		// 	},
		// }

		// check max range
		currentBlock = 27266647
		if currentBlock-blockStart > 5000 {
			blockStart = currentBlock - 5000
		}
		fmt.Println(currentBlock) // 5671744

		query := &bind.FilterOpts{
			Start: blockStart,
			End:   &currentBlock,
		}
		// logs, err := client.FilterLogs(context.Background(), query)
		// if err != nil {
		// 	return err
		// }

		eventsIterator, err := contract.FilterRewardReferral(query, nil, nil, nil)
		if err != nil {
			logger.Error().Msg("Get FilterRewardReferral failed")
			return err
		}

		m.parseRewardReferralEvent(eventsIterator)
	}

	// eg := new(errgroup.Group)
	// eg.Go(func() error {
	// 	return m.watchRewardReferralEvent(ctx, contract)
	// })
	// eg.Go(func() error {
	// 	return m.watchOwnershipTransferred(ctx, contract)
	// })
	// if err = eg.Wait(); err != nil {
	// 	logger.Error().Msg("Events error group throw error")
	// 	return err
	// }
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
				RewardReferralEvent{
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
		j, _ := json.MarshalIndent(
			RewardReferralEvent{
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

	if eventsIterator.Error() != nil {
		return eventsIterator.Error()
	}

	return nil
}
