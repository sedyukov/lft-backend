package main

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
	"github.com/sedyukov/lft-backend/internal/blockchain"
	lftdb "github.com/sedyukov/lft-backend/internal/database/lft"
	"github.com/sedyukov/lft-backend/internal/service"
	"github.com/spf13/viper"
)

func main() {
	// Load viper config
	err := service.LoadConfig()
	if err != nil {
		panic(err)
	}

	// Start logger service
	logger, err := service.NewLogger("parser", service.LoggerConfig{
		OutOnly: true,
	})
	if err != nil {
		panic(err)
	}
	logger.Info().Msg("Logger sucessfully started")

	// Initialize database and migrate
	var migrateDatabase = true
	err = lftdb.InitDatabase(logger, migrateDatabase)
	if err != nil {
		panic(err)
	}
	logger.Info().Msg("DB init sucessfully")

	establishRpcMonitoring(logger)
}

func establishWsMonitoring(logger zerolog.Logger) {
	var (
		bscWs           = viper.GetString("ENDPOINT_WS")
		contractAddress = viper.GetString("CONTRACT_ADDRESS")
	)

	// Init client
	ctx, cancel := context.WithCancel(context.Background())
	client, err := ethclient.DialContext(ctx, bscWs)
	if err != nil {
		logger.Error().Msg("Connection failed to: " + bscWs)
		panic(err)
	}

	nId, err := client.NetworkID(ctx)
	if err != nil {
		logger.Error().Msg("Connection failed to: " + bscWs)
		panic(err)
	}
	logger.Info().Msg(nId.String())
	defer cancel()
	logger.Info().Msg("Client init sucessfully")

	// Start monitoring
	monitor := blockchain.NewMonitor(contractAddress, logger)
	err = monitor.Start(ctx, client, logger)
	if err != nil {
		panic(err)
	}
}

func establishRpcMonitoring(logger zerolog.Logger) {
	var (
		rpcEndpoint     = viper.GetString("ENDPOINT_RPC")
		contractAddress = viper.GetString("CONTRACT_ADDRESS")
	)

	// Init client
	ctx, cancel := context.WithCancel(context.Background())
	client, err := ethclient.DialContext(ctx, rpcEndpoint)
	if err != nil {
		logger.Error().Msg("Connection failed to: " + rpcEndpoint)
		panic(err)
	}

	nId, err := client.NetworkID(ctx)
	if err != nil {
		logger.Error().Msg("Connection failed to: " + rpcEndpoint)
		panic(err)
	}
	logger.Info().Msg(nId.String())
	defer cancel()
	logger.Info().Msg("Client init sucessfully")

	// Start monitoring
	monitor := blockchain.NewMonitor(contractAddress, logger)
	err = monitor.StartRpc(ctx, client, logger)
	if err != nil {
		panic(err)
	}
}
