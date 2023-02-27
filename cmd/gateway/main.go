package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"

	lftdb "github.com/sedyukov/lft-backend/internal/database/lft"
	"github.com/sedyukov/lft-backend/internal/routes"
	"github.com/sedyukov/lft-backend/internal/service"
)

func main() {
	// Load viper config
	err := service.LoadConfig()
	if err != nil {
		panic(err)
	}

	// Start logger service
	logger, err := service.NewLogger("gateway", service.LoggerConfig{
		OutOnly: true,
	})
	if err != nil {
		panic(err)
	}
	logger.Info().Msg("Logger sucessfully started for gateway")
	app := fiber.New()

	// Initialize database without migration
	var migrateDatabase = false
	err = lftdb.InitDatabase(logger, migrateDatabase)
	if err != nil {
		panic(err)
	}
	logger.Info().Msg("DB init finished")

	// Setup gateway routes
	routes.SetupGatewayRoutes(app)

	// Listening for requests
	var port = viper.GetString("GATEWAY_PORT")
	logger.Info().Msgf("Listening to port %v", port)
	app.Listen(":" + port)
}
