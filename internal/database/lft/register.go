package lftdb

import (
	"gorm.io/gorm"
)

type Register struct {
	gorm.Model
	Refferal    string `json:"refferal"`
	Trader      string `json:"trader"`
	BlockHeight int64  `json:"block_height"`
}

func GetAllRegister() []Register {
	db := DBInstance.con
	var rs []Register
	db.Find(&rs)
	return rs
}

func GetRegister(id string) Register {
	db := DBInstance.con
	var r Register
	db.First(&r, id)
	return r
}
