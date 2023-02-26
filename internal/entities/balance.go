package entities

type Balance struct {
	ID             int64  `gorm:"column:id"`
	AccountAddress string `gorm:"column:accountAddress"`
	Amount         string `gorm:"column:amount"`
	TokenAddress   string `gorm:"column:tokenAddress"`
}
