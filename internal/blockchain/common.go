package blockchain

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	contracts "github.com/sedyukov/lft-backend/contracts/interfaces"
)

var (
	ErrInvalidKey             = errors.New("invalid key")
	ErrInvalidAddress         = errors.New("invalid address")
	ErrInvalidContractAddress = errors.New("invalid contract address")
)

type SignerConfig struct {
	privKeyHex string `mapstructure:"priv_key_hex"`
	GasLimit   int64  `mapstructure:"gas_limit"`
	GasPrice   int64  `mapstructure:"gas_price"`
	WeiFounds  int64  `mapstructure:"default_wei_founds"`
}

// getSigner get the signer for sign transactions
func getSigner(ctx context.Context, client *ethclient.Client, config SignerConfig) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(config.privKeyHex)
	if err != nil {
		return nil, err
	}
	publicKey, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, ErrInvalidKey
	}

	address := crypto.PubkeyToAddress(*publicKey)
	nonce, err := client.PendingNonceAt(ctx, address)
	if err != nil {
		return nil, err
	}
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	signer, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	signer.Nonce = big.NewInt(int64(nonce))
	signer.Value = big.NewInt(config.WeiFounds)
	signer.GasLimit = uint64(config.GasLimit)
	signer.GasPrice = big.NewInt(config.GasPrice)

	return signer, nil
}

// getContract get an instance of the deployed contract
func getContract(ctx context.Context, client *ethclient.Client, contractAddress string) (*contracts.Contract, error) {
	err := validateContractAddress(ctx, client, contractAddress)
	if err != nil {
		return nil, err
	}
	contract, err := contracts.NewContract(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, err
	}
	return contract, nil
}

// validateContractAddress validate the contract address checking if the contract is deployed
func validateContractAddress(ctx context.Context, client *ethclient.Client, address string) error {
	if err := validateAddress(address); err != nil {
		return err
	}
	contractAddress := common.HexToAddress(address)
	bytecode, err := client.CodeAt(ctx, contractAddress, nil)
	if err != nil {
		return err
	}
	if len(bytecode) > 0 {
		return nil
	}
	return ErrInvalidContractAddress
}

// validateAddress validate address format
func validateAddress(address string) error {
	regex := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if ok := regex.MatchString(address); !ok {
		return ErrInvalidAddress
	}
	return nil
}

// etherToWei convert Ether to Wei
func etherToWei(eth *big.Int) *big.Int {
	return new(big.Int).Mul(eth, big.NewInt(params.Ether))
}

// weiToEther convert Wei to Ether
func weiToEther(wei *big.Int) *big.Int {
	return new(big.Int).Div(wei, big.NewInt(params.Ether))
}

// DurationToString converts provided duration to human readable string presentation.
func DurationToString(d time.Duration) string {
	ns := time.Duration(d.Nanoseconds())
	ms := float64(ns) / 1000000.0
	var unit string
	var amount string
	if ns < time.Microsecond {
		amount, unit = humanize.CommafWithDigits(float64(ns), 0), "ns"
	} else if ns < time.Millisecond {
		amount, unit = humanize.CommafWithDigits(ms*1000.0, 3), "μs"
	} else if ns < time.Second {
		amount, unit = humanize.CommafWithDigits(ms, 3), "ms"
	} else if ns < time.Minute {
		amount, unit = humanize.CommafWithDigits(ms/1000.0, 3), "s"
	} else if ns < time.Hour {
		amount, unit = humanize.CommafWithDigits(ms/60000.0, 3), "m"
	} else if ns < 24*time.Hour {
		amount, unit = humanize.CommafWithDigits(ms/3600000.0, 3), "h"
	} else {
		days := ms / 86400000.0
		unit = "day"
		if days > 1 {
			unit = "days"
		}
		amount = humanize.CommafWithDigits(days, 3)
	}
	return fmt.Sprintf("%s %s", amount, unit)
}

// func pringLogs() {
// 	for _, vLog := range logs {
// 		event := struct {
// 		  Key   [32]byte
// 		  Value [32]byte
// 		}{}
// 		err := contractAbi.Unpack(&event, "ItemSet", vLog.Data)
// 		if err != nil {
// 		  log.Fatal(err)
// 		}

// 		fmt.Println(string(event.Key[:]))   // foo
// 		fmt.Println(string(event.Value[:])) // bar
// 	  }
// }
