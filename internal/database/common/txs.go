package parser

import "github.com/sedyukov/lft-backend/internal/entities"

func (db *DB) GetTxByHash(hash int64) (entities.Tx, error) {
	var record entities.Tx

	err := db.con.Table("Txes").Take(&record, "hash = ?", hash).Error
	if err != nil {
		return entities.Tx{}, err
	}

	return record, nil
}

func (db *DB) GetLastTx() (entities.Tx, error) {
	var record entities.Tx

	err := db.con.Table("Txes").Order("\"blockHeight\" DESC").Take(&record).Error
	if err != nil {
		return entities.Tx{}, err
	}

	return record, nil
}
