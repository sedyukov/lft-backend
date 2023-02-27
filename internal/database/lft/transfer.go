package lftdb

import (
	"gorm.io/gorm"
)

type Transfer struct {
	gorm.Model
	From        string `json:"from"`
	To          string `json:"to"`
	Value       string `json:"value"`
	BlockHeight int64  `json:"blockHeight"`
}

func CreateTransfer() {

}
