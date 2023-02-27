package lftdb

import (
	"gorm.io/gorm"
)

type Stake struct {
	gorm.Model
	Staker      string `json:"staker"`
	Amount      string `json:"amount"`
	BlockHeight int64  `json:"block_height"`
}

func CreateStake() {

}
