package lftcontrollers

import (
	"github.com/gofiber/fiber/v2"

	lftdb "github.com/sedyukov/lft-backend/internal/database/lft"
)

func GetAllRegister(c *fiber.Ctx) error {
	var rs = lftdb.GetAllRegister()
	c.JSON(rs)
	return nil
}

func GetRegister(c *fiber.Ctx) error {
	id := c.Params("id")
	var r = lftdb.GetRegister(id)
	c.JSON(r)
	return nil
}
