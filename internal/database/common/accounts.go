package parser

import "github.com/sedyukov/lft-backend/internal/entities"

func (db *DB) GetAccountByAddress(accountAddress string) (entities.Account, error) {
	var address entities.Account

	err := db.con.Table("Accounts").Take(&address, `"address" = ?`, accountAddress).Error
	if err != nil {
		return entities.Account{}, err
	}

	return address, nil
}
