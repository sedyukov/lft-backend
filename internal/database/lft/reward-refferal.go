package lftdb

import (
	"gorm.io/gorm"
)

type RewardReferral struct {
	gorm.Model
	Trader      string `json:"trader"`
	Refferal    string `json:"refferal"`
	Level       uint8  `json:"level"`
	Amount      string `json:"amount"`
	BlockNumber uint64 `json:"block_number"`
}

type RewardSumResult struct {
	sum string
}

type RewardSumLevelsResult struct {
	Sum   string `json:"sum"`
	Level uint8  `json:"level"`
	Count uint64 `json:"count"`
}

func GetSumRewardsByRefAddress(refferal string) string {
	db := DBInstance.con
	var sum string
	sql := "select sum(amount::numeric) from reward_referrals rr where refferal = ?"
	db.Raw(sql, refferal).Scan(&sum)
	return sum
}

func GetSumRewardsByRefAddressAndLevels(refferal string) []RewardSumLevelsResult {
	db := DBInstance.con
	var res []RewardSumLevelsResult
	sql := "select level, sum(amount::numeric), count(amount) from reward_referrals rr where refferal = ? group by level"
	db.Raw(sql, refferal).Scan(&res)
	return res
}

func GetRewardReferral(id string) RewardReferral {
	db := DBInstance.con
	var rr RewardReferral
	db.First(&rr, id)
	return rr
}

func GetAllRewardReferral() []RewardReferral {
	db := DBInstance.con
	var rrs []RewardReferral
	db.Find(&rrs)
	return rrs
}

func CreateRewardRefferal(rr RewardReferral) {
	db := DBInstance.con
	db.Create(&rr)
}
