package lftdb

import (
	"gorm.io/gorm"
)

type Unstake struct {
	gorm.Model
	Staker      string `json:"staker"`
	Amount      string `json:"amount"`
	BlockHeight int64  `json:"blockHeight"`
}

func CreateUnstake() {

}
