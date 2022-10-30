package dispatch

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/khafidprayoga/psychic-octo-disco/domain"
	request "github.com/khafidprayoga/psychic-octo-disco/http/req"
	"github.com/khafidprayoga/psychic-octo-disco/utils"
	"strconv"
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
		todoId := ctx.Params("id")

		httpCode, errLogic, _ := h.useCaseImpl.DeleteExistingTodo(todoId)
		if errLogic != nil {
			return ctx.Status(httpCode).JSON(
				utils.ErrorResponse(errLogic, 4001),
			)
		}

		successMsg := fmt.Sprintf("deleted todo with id %v", todoId)
		return ctx.Status(httpCode).JSON(
			utils.SuccessResponse(successMsg, nil),
		)
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

		resData, httpCode, errLogic, internalErr := h.useCaseImpl.CreateNewTodo(req)
		if errLogic != nil {
			return ctx.Status(httpCode).JSON(
				utils.ErrorResponse(errLogic, internalErr),
			)
		}

		return ctx.Status(httpCode).JSON(
			utils.SuccessResponse("created new todo", &resData),
		)
	}
}

func (h *TodoHandler) UpdateData() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req request.UpdateExistingTodo

		todoId := ctx.Params("id")
		idInt, errConv := strconv.Atoi(todoId)
		if errConv != nil {
			err := errors.New("invalid todo id")
			return ctx.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse(err, utils.HTTPRequestErr))
		}
		req.SetId(idInt)

		if parsingErr := utils.JSONBodyParser(ctx, &req); parsingErr != nil {
			panic(parsingErr)
			return parsingErr
		}

		resData, httpCode, errLogic, internalErr := h.useCaseImpl.UpdateExistingTodo(req)
		if errLogic != nil {
			return ctx.Status(httpCode).JSON(
				utils.ErrorResponse(errLogic, internalErr),
			)
		}

		return ctx.Status(httpCode).JSON(
			utils.SuccessResponse("succes update existing todo", &resData),
		)
	}
}
