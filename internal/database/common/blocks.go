package parser

import "github.com/sedyukov/lft-backend/internal/entities"

func (db *DB) GetBlockByHeight(height int64) (entities.Block, error) {
	var block entities.Block

	err := db.con.Table("Blocks").Take(&block, "height = ?", height).Error
	if err != nil {
		return entities.Block{}, err
	}

	return block, nil
}

func (db *DB) GetLastBlock() (entities.Block, error) {
	var block entities.Block

	err := db.con.Table("Blocks").Order("height DESC").Take(&block).Error
	if err != nil {
		return entities.Block{}, err
	}

	return block, nil
}
