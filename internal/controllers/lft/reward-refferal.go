package lftcontrollers

import (
	"math/big"

	"github.com/gofiber/fiber/v2"

	lftdb "github.com/sedyukov/lft-backend/internal/database/lft"
)

type RewardReferralEvent struct {
	Trader      string   `json:"trader"`
	Refferal    string   `json:"refferal"`
	Level       uint8    `json:"level"`
	Amount      *big.Int `json:"amount"`
	BlockNumber uint64   `json:"block_number"`
}

type RewardRefferalSumResponse struct {
	Referral string `json:"refferal"`
	Sum      string `json:"amount"`
}

type RewardsRefferalSumWithLevelsResponse struct {
	Referral string                        `json:"refferal"`
	Rewards  []lftdb.RewardSumLevelsResult `json:"rewards"`
}

func GetAllRewardReferral(c *fiber.Ctx) error {
	var rrs = lftdb.GetAllRewardReferral()
	c.JSON(rrs)
	return nil
}

func GetSumRewardsByRefAddress(c *fiber.Ctx) error {
	address := c.Params("address")
	sum := lftdb.GetSumRewardsByRefAddress(address)

	res := RewardRefferalSumResponse{
		Referral: address,
		Sum:      sum,
	}

	c.JSON(res)
	return nil
}

func GetSumRewardsByRefAddressWithLevels(c *fiber.Ctx) error {
	address := c.Params("address")
	dbRes := lftdb.GetSumRewardsByRefAddressAndLevels(address)

	res := RewardsRefferalSumWithLevelsResponse{
		Referral: address,
		Rewards:  dbRes,
	}

	c.JSON(res)
	return nil
}

func GetRewardReferral(c *fiber.Ctx) error {
	id := c.Params("id")
	var rr = lftdb.GetRewardReferral(id)
	c.JSON(rr)
	return nil
}

func CreateRewardRefferal(rre RewardReferralEvent) {
	rr := lftdb.RewardReferral{
		Trader:      rre.Trader,
		Refferal:    rre.Refferal,
		Level:       rre.Level,
		Amount:      rre.Amount.String(),
		BlockNumber: rre.BlockNumber,
	}
	lftdb.CreateRewardRefferal(rr)
}
