package dispatch

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khafidprayoga/psychic-octo-disco/domain"
)

func Routes(app *fiber.App, handler domain.ActivityHandler) {
	activityEndpoint := app.Group("/activity-groups")

	activityEndpoint.Delete("/:id", handler.DeleteExistingActivity()).Name("delete activity by id")
	activityEndpoint.Get("/:id", handler.GetDetailActivity()).Name("get detail existing activity")
	activityEndpoint.Get("", handler.GetAllActivity()).Name("get collections of existing activity")
	activityEndpoint.Post("", handler.PostNewActivity()).Name("create new one activity")
}
