package lftdb

import (
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

var (
	DBInstance *DB
)

func InitDatabase(logger zerolog.Logger, autoMigrate bool) error {
	db, err := NewDB(Config{
		Host:     viper.GetString("PSQL_PARSER_HOST"),
		User:     viper.GetString("PSQL_PARSER_USER"),
		Pass:     viper.GetString("PSQL_PARSER_PASS"),
		Port:     viper.GetUint32("PSQL_PARSER_PORT"),
		Database: viper.GetString("PSQL_PARSER_DB"),
	}, logger)

	if err != nil {
		return err
	}

	if autoMigrate {
		logger.Info().Msg("DB migration started")
		err := db.con.AutoMigrate(
			&OwnershipTransferred{},
			&Register{},
			&RewardReferral{},
			&RewardStakers{},
			&Stake{},
			&Transfer{},
			&Unstake{},
			&Counter{},
		)
		if err != nil {
			return err
		}
		logger.Info().Msg("DB migration finished")
	}
	DBInstance = db

	return nil
}
