package dispatch

import (
	"fmt"
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
		todoId := ctx.Params("id")

		httpCode, errLogic, internalErr := h.useCaseImpl.DeleteExistingTodo(todoId)
		if errLogic != nil {
			// FOR TEST
			if httpCode == fiber.StatusNotFound {
				return ctx.Status(httpCode).JSON(
					fiber.Map{
						"status":  "Not Found",
						"message": fmt.Sprintf("Todo with ID %v Not Found", todoId),
						"data":    struct{}{},
					},
				)
			}

			return ctx.Status(httpCode).JSON(
				utils.ErrorResponse(errLogic, internalErr),
			)
		}

		return ctx.Status(httpCode).JSON(
			utils.SuccessResponse("Success", struct{}{}),
		)
	}
}

func (h *TodoHandler) GetDetail() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		todoId := ctx.Params("id")

		resData, httpCode, errLogic, internalErr := h.useCaseImpl.GetDetailTodo(todoId)
		if errLogic != nil {
			// FOR TEST
			if httpCode == fiber.StatusNotFound {
				return ctx.Status(httpCode).JSON(
					fiber.Map{
						"status":  "Not Found",
						"message": fmt.Sprintf("Todo with ID %v Not Found", todoId),
						"data":    struct{}{},
					},
				)
			}
			return ctx.Status(httpCode).JSON(
				utils.ErrorResponse(errLogic, internalErr),
			)
		}

		return ctx.Status(httpCode).JSON(
			utils.SuccessResponse("success get detail todo by id", &resData),
		)
	}
}

func (h *TodoHandler) GetList() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		activityGroupId := ctx.Query("activity_group_id")
		resData, httpCode, errLogic, internalErr := h.useCaseImpl.GetAllListTodo(activityGroupId)
		if errLogic != nil {
			return ctx.Status(httpCode).JSON(
				utils.ErrorResponse(errLogic, internalErr),
			)
		}

		return ctx.Status(httpCode).JSON(
			utils.SuccessResponse("succes get list todo", &resData),
		)
	}
}

func (h *TodoHandler) PostNewData() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req request.CreateNewTodo

		if parsingErr := utils.JSONBodyParser(ctx, &req); parsingErr != nil {
			return parsingErr
		}

		// FOR TESTING
		if req.Title == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON(
				fiber.Map{
					"status":  "Bad Request",
					"message": "title cannot be null",
				},
			)
		}

		// FOR TESTING
		if req.ActivityGroupID == nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(
				fiber.Map{
					"status":  "Bad Request",
					"message": "activity_group_id cannot be null",
				},
			)
		}
		resData, httpCode, errLogic, internalErr := h.useCaseImpl.CreateNewTodo(req)
		if errLogic != nil {
			return ctx.Status(httpCode).JSON(
				utils.ErrorResponse(errLogic, internalErr),
			)
		}

		return ctx.Status(httpCode).JSON(
			utils.SuccessResponse("created new todo", resData.ToResponse()),
		)
	}
}

func (h *TodoHandler) UpdateData() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req request.UpdateExistingTodo

		todoId := ctx.Params("id")
		req.SetId(todoId)

		if parsingErr := utils.JSONBodyParser(ctx, &req); parsingErr != nil {
			return parsingErr
		}

		resData, httpCode, errLogic, internalErr := h.useCaseImpl.UpdateExistingTodo(req)
		if errLogic != nil {
			// FOR TEST
			if httpCode == fiber.StatusNotFound {
				return ctx.Status(httpCode).JSON(
					fiber.Map{
						"status":  "Not Found",
						"message": fmt.Sprintf("Todo with ID %v Not Found", todoId),
						"data":    struct{}{},
					},
				)
			}
			return ctx.Status(httpCode).JSON(
				utils.ErrorResponse(errLogic, internalErr),
			)
		}

		return ctx.Status(httpCode).JSON(
			utils.SuccessResponse("succes update existing todo", &resData),
		)
	}
}
