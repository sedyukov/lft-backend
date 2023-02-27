package lftcontrollers

import (
	"github.com/gofiber/fiber/v2"

	lftdb "github.com/sedyukov/lft-backend/internal/database/lft"
)

func GetAllRewardStakers(c *fiber.Ctx) error {
	var rss = lftdb.GetAllRewardStakers()
	c.JSON(rss)
	return nil
}

func GetRewardStakers(c *fiber.Ctx) error {
	id := c.Params("id")
	var rs = lftdb.GetRewardStakers(id)
	c.JSON(rs)
	return nil
}
