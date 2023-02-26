package main

import (
	"github.com/sedyukov/lft-backend/internal/service"
)

func main() {
	err := service.LoadConfig()
	if err != nil {
		panic(err)
	}

	logger, err := service.NewLogger("parser", service.LoggerConfig{
		OutOnly: true,
	})
	if err != nil {
		panic(err)
	}
	logger.Info().Msg("Logger sucessfully started for gateway")

}
