package lftdb

import (
	"gorm.io/gorm"
)

type OwnershipTransferred struct {
	gorm.Model
	OldOwner    string `json:"old_owner"`
	NewOwner    string `json:"new_owner"`
	BlockHeight int64  `json:"block_height"`
}

func GetAllOwnershipTransferred() []OwnershipTransferred {
	db := DBInstance.con
	var ots []OwnershipTransferred
	db.Find(&ots)
	return ots
}

func GetOwnershipTransferred(id string) OwnershipTransferred {
	db := DBInstance.con
	var ot OwnershipTransferred
	db.First(&ot, id)
	return ot
}
