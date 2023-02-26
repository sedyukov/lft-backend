package entities

import "database/sql"

type Token struct {
	address         sql.NullString `gorm:"column:address"`
	ContractAddress sql.NullString `gorm:"column:contractAddress"`
}
