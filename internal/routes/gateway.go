package routes

import (
	"github.com/gofiber/fiber/v2"

	lftcontrollers "github.com/sedyukov/lft-backend/internal/controllers/lft"
)

func SetupGatewayRoutes(app *fiber.App) {
	// ownership transferred
	app.Get("/api/v1/ownership-transferred", lftcontrollers.GetAllOwnershipTransferred)
	app.Get("/api/v1/ownership-transferred/:id", lftcontrollers.GetOwnershipTransferred)

	// register
	app.Get("/api/v1/register", lftcontrollers.GetAllRegister)
	app.Get("/api/v1/register/:id", lftcontrollers.GetRegister)

	// reward refferal
	app.Get("/api/v1/reward-refferal", lftcontrollers.GetAllRewardReferral)
	app.Get("/api/v1/reward-refferal/:id", lftcontrollers.GetRewardReferral)
	app.Get("/api/v1/rewards-sum/ref/:address", lftcontrollers.GetSumRewardsByRefAddress)

	// reward stakers
	app.Get("/api/v1/reward-stakers", lftcontrollers.GetAllRewardStakers)
	app.Get("/api/v1/reward-stakers/:id", lftcontrollers.GetRewardStakers)
}
