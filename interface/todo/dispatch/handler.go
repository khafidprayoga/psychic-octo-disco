package dispatch

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khafidprayoga/psychic-octo-disco/domain"
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
		return nil
	}
}

func (h *TodoHandler) UpdateData() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return nil
	}
}
