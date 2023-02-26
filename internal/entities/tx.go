package entities

import "database/sql"

type Tx struct {
	Hash        string         `gorm:"column:hash"`
	From        sql.NullString `gorm:"column:from"`
	To          sql.NullString `gorm:"column:to"`
	Value       sql.NullString `gorm:"column:value"`
	BlockHeight int64          `gorm:"column:blockHeight"`
	GasUsed     int64          `gorm:"column:gasUsed"`
	GasPrice    int64          `gorm:"column:gasPrice"`
	Nonce       sql.NullInt64  `gorm:"column:nonce"`
}
