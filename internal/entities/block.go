package entities

type Block struct {
	Height        int64  `gorm:"column:height"`
	Timestamp     string `gorm:"column:date"`
	Hash          string `gorm:"column:hash"`
	Miner         string `gorm:"column:miner"`
	BaseFeePerGas string `gorm:"column:baseFeePerGas"`
	GasUsed       string `gorm:"column:gasUsed"`
	GasLimit      string `gorm:"column:gasLimit"`
	TxsCount      int64  `gorm:"column:txsCount"`
}
