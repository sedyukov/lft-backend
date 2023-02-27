package lftdb

import (
	"gorm.io/gorm"
)

type RewardReferral struct {
	gorm.Model
	Trader      string `json:"trader"`
	Refferal    string `json:"refferal"`
	Level       string `json:"level"`
	Amount      string `json:"amount"`
	BlockHeight int64  `json:"block_height"`
}

func GetAllRewardReferral() []RewardReferral {
	db := DBInstance.con
	var rrs []RewardReferral
	db.Find(&rrs)
	return rrs
}

func GetRewardReferral(id string) RewardReferral {
	db := DBInstance.con
	var rr RewardReferral
	db.First(&rr, id)
	return rr
}
