package lftcontrollers

import (
	"github.com/gofiber/fiber/v2"

	lftdb "github.com/sedyukov/lft-backend/internal/database/lft"
)

func GetAllOwnershipTransferred(c *fiber.Ctx) error {
	var ots = lftdb.GetAllOwnershipTransferred()
	c.JSON(ots)
	return nil
}

func GetOwnershipTransferred(c *fiber.Ctx) error {
	id := c.Params("id")
	var ot = lftdb.GetOwnershipTransferred(id)
	c.JSON(ot)
	return nil
}
