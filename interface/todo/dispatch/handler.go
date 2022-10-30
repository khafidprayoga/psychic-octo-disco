package dispatch

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khafidprayoga/psychic-octo-disco/domain"
	request "github.com/khafidprayoga/psychic-octo-disco/http/req"
	"github.com/khafidprayoga/psychic-octo-disco/utils"
)

type TodoHandler struct {
	useCaseImpl domain.TodoUseCase
	dataImpl    domain.TodoData
}

func New(uc domain.TodoUseCase, d domain.TodoData) domain.TodoHandler {
	return &TodoHandler{
		useCaseImpl: uc,
		dataImpl:    d,
	}
}

func (h *TodoHandler) DeleteData() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return nil
	}
}

func (h *TodoHandler) GetDetail() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return nil
	}
}

func (h *TodoHandler) GetList() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return nil
	}
}

func (h *TodoHandler) PostNewData() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req request.CreateNewTodo

		if parsingErr := utils.JSONBodyParser(ctx, &req); parsingErr != nil {
			return parsingErr
		}

		resData, httpCode, errLogic, errType := h.useCaseImpl.CreateNewTodo(req)
		if errLogic != nil {
			return ctx.Status(httpCode).JSON(
				utils.ErrorResponse(errType, errLogic),
			)
		}

		return ctx.Status(httpCode).JSON(
			utils.SuccessResponse("created new todo", &resData),
		)
	}
}

func (h *TodoHandler) UpdateData() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return nil
	}
}
