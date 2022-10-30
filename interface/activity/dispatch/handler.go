package dispatch

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/khafidprayoga/psychic-octo-disco/domain"
	request "github.com/khafidprayoga/psychic-octo-disco/http/req"
	"github.com/khafidprayoga/psychic-octo-disco/utils"
)

type ActivityHandler struct {
	useCaseImpl domain.ActivityUseCase
	dataImpl    domain.ActivityData
}

func New(uc domain.ActivityUseCase, data domain.ActivityData) domain.ActivityHandler {
	return &ActivityHandler{
		useCaseImpl: uc,
		dataImpl:    data,
	}
}

func (h *ActivityHandler) DeleteExistingActivity() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		activityId := ctx.Params("id")

		httpCode, errLogic, internalErr := h.useCaseImpl.DeleteActivity(activityId)
		if errLogic != nil {
			return ctx.Status(httpCode).JSON(
				utils.ErrorResponse(errLogic, internalErr),
			)
		}

		successMsg := fmt.Sprintf("deleted activity with id %v", activityId)
		return ctx.Status(httpCode).JSON(
			utils.SuccessResponse(successMsg, nil),
		)
	}
}

func (h *ActivityHandler) GetDetailActivity() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		idActivity := ctx.Params("id")

		resData, httpCode, errLogic, internalErr := h.useCaseImpl.DetailActivity(idActivity)
		if errLogic != nil {
			//FOR TEST
			if httpCode == fiber.StatusNotFound {
				return ctx.Status(httpCode).JSON(
					fiber.Map{
						"status":  "Not Found",
						"message": fmt.Sprintf("Activity with ID %v Not Found", idActivity),
						"data":    struct{}{},
					},
				)
			}

			return ctx.Status(httpCode).JSON(
				utils.ErrorResponse(errLogic, internalErr),
			)
		}

		return ctx.Status(httpCode).JSON(
			utils.SuccessResponse("success get detail activity by id", &resData),
		)
	}
}

func (h *ActivityHandler) GetAllActivity() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		resData, httpCode, errLogic, internalErr := h.useCaseImpl.ListActivity()
		if errLogic != nil {
			return ctx.Status(httpCode).JSON(
				utils.ErrorResponse(errLogic, internalErr),
			)
		}

		return ctx.Status(httpCode).JSON(
			utils.SuccessResponse("success get list activity", &resData),
		)
	}
}

func (h *ActivityHandler) PostNewActivity() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req request.CreateNewActivity

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

		resData, httpCode, errLogic, internalErr := h.useCaseImpl.CreateNewActivity(req)
		if errLogic != nil {
			return ctx.Status(httpCode).JSON(
				utils.ErrorResponse(errLogic, internalErr),
			)
		}

		return ctx.Status(httpCode).JSON(
			utils.SuccessResponse("created new activity", &resData),
		)
	}
}

func (h *ActivityHandler) PatchExistingActivity() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req request.UpdateExistingActivity

		activityId := ctx.Params("id")
		req.SetId(activityId)

		if parsingErr := utils.JSONBodyParser(ctx, &req); parsingErr != nil {
			return parsingErr
		}

		resData, httpCode, errLogic, internalErr := h.useCaseImpl.UpdateActivity(req)
		if errLogic != nil {
			//FOR TEST
			if httpCode == fiber.StatusNotFound {
				return ctx.Status(httpCode).JSON(
					fiber.Map{
						"status":  "Not Found",
						"message": fmt.Sprintf("Activity with ID %v Not Found", activityId),
						"data":    struct{}{},
					},
				)
			}
			return ctx.Status(httpCode).JSON(
				utils.ErrorResponse(errLogic, internalErr),
			)
		}

		return ctx.Status(httpCode).JSON(
			utils.SuccessResponse("success update existing activity", &resData),
		)
	}
}
