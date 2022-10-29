package dispatch

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khafidprayoga/psychic-octo-disco/domain"
)

func Routes(app *fiber.App, handler domain.TodoHandler) {
	todoEndpoint := app.Group("/todo-items")

	todoEndpoint.Delete("/:id", handler.DeleteData()).Name("delete single todo")
	todoEndpoint.Get("", handler.GetList()).Name("get list collections of todo")
	todoEndpoint.Get("/:id", handler.GetDetail()).Name("get detail single todo")
	todoEndpoint.Patch("/:id", handler.UpdateData()).Name("update single todo")
	todoEndpoint.Post("", handler.PostNewData()).Name("insert a new todo")

}
