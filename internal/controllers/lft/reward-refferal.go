package lftcontrollers

import (
	"github.com/gofiber/fiber/v2"

	lftdb "github.com/sedyukov/lft-backend/internal/database/lft"
)

func GetAllRewardReferral(c *fiber.Ctx) error {
	var rrs = lftdb.GetAllRewardReferral()
	c.JSON(rrs)
	return nil
}

func GetRewardReferral(c *fiber.Ctx) error {
	id := c.Params("id")
	var rr = lftdb.GetRewardReferral(id)
	c.JSON(rr)
	return nil
}
