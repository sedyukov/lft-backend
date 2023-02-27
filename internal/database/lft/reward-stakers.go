package lftdb

import (
	"gorm.io/gorm"
)

type RewardStakers struct {
	gorm.Model
	Trader      string `json:"trader"`
	Amount      string `json:"amount"`
	BlockHeight int64  `json:"block_height"`
}

func GetAllRewardStakers() []RewardStakers {
	db := DBInstance.con
	var rss []RewardStakers
	db.Find(&rss)
	return rss
}

func GetRewardStakers(id string) RewardStakers {
	db := DBInstance.con
	var rs RewardStakers
	db.First(&rs, id)
	return rs
}
