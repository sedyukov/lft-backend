package lftdb

import (
	"gorm.io/gorm"
)

type Counter struct {
	gorm.Model
	Key   string `json:"key"`
	Value string `json:"value"`
}

func CreateCounter(c Counter) {
	db := DBInstance.con
	db.Create(&c)
}

func GetLastBlock() string {
	db := DBInstance.con
	var res Counter

	db.Table("counters").Select("value").Where("key = ?", "block").Scan(&res)
	if res.Value == "" {
		panic("Start block not defined")
	}

	return res.Value
}

func UpdateLastBlock(block string) {
	db := DBInstance.con

	db.Table("counters").Where("key = ?", "block").Update("value", block)
}
