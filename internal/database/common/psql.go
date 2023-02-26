package parser

import (
	"database/sql"
	"fmt"

	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	con    *gorm.DB
	sqlDB  *sql.DB
	logger zerolog.Logger
}

type Config struct {
	Host     string
	User     string
	Pass     string
	Database string
	Port     uint32
}

func NewDB(config Config, logger zerolog.Logger) (*DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		config.Host,
		config.User,
		config.Pass,
		config.Database,
		config.Port,
	)
	con, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db, err := con.DB()
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DB{
		con:    con,
		sqlDB:  db,
		logger: logger,
	}, nil
}

func (db *DB) Close() error {
	return db.sqlDB.Close()
}
