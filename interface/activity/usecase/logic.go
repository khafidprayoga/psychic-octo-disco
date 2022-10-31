package usecase

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/khafidprayoga/psychic-octo-disco/domain"
	"github.com/khafidprayoga/psychic-octo-disco/http/entities"
	request "github.com/khafidprayoga/psychic-octo-disco/http/req"
	"github.com/khafidprayoga/psychic-octo-disco/interface/activity/interfaceError"
	"github.com/khafidprayoga/psychic-octo-disco/utils"
	"log"
)

type ActivityUseCase struct {
	data domain.ActivityData
}

func New(data domain.ActivityData) *ActivityUseCase {
	return &ActivityUseCase{
		data,
	}
}

func (u *ActivityUseCase) CreateNewActivity(req request.CreateNewActivity) (res *entities.Activity, httpCode int, errType error, srvError int) {
	if err := utils.ValidateStruct[request.CreateNewActivity](req); err != nil {
		errMsg := fmt.Errorf("%v, reason: %v", interfaceError.FailedCreateNewActivity, err)
		return nil, fiber.StatusBadRequest, errMsg, utils.HTTPRequestErr
	}

	resData, err := u.data.CreateNewActivity(req)
	if err != nil {
		return nil, fiber.StatusInternalServerError, err, utils.DatabaseError
	}

	return resData, fiber.StatusCreated, nil, 0
}

func (u *ActivityUseCase) DetailActivity(id string) (res *entities.Activity, httpCode int, errType error, srvError int) {
	// Validate activity exist
	if err := u.data.ValidateActivity(id); err != nil {
		return nil, fiber.StatusNotFound, interfaceError.DataNotFound, utils.HTTPRequestErr
	}

	entitiesData, err := u.data.DetailActivityData(id)
	if err != nil {
		return nil, fiber.StatusInternalServerError, err, utils.DatabaseError
	}

	return entitiesData, fiber.StatusOK, nil, 0
}

func (u *ActivityUseCase) DeleteActivity(id string) (httpCode int, errType error, srvError int) {
	// Validate activity exist
	if err := u.data.ValidateActivity(id); err != nil {
		return fiber.StatusNotFound, err, utils.NotFoundErr
	}

	if err := u.data.DeleteActivityData(id); err != nil {
		log.Println(err)
		return fiber.StatusInternalServerError, err, utils.DatabaseError

	}

	return fiber.StatusOK, nil, 0

}

func (u *ActivityUseCase) ListActivity() (res []entities.Activity, httpCode int, errType error, srvError int) {
	listData, err := u.data.ListActivityData()
	if err != nil {
		return res, fiber.StatusInternalServerError, err, utils.DatabaseError
	}

	return listData, fiber.StatusOK, nil, 0
}

func (u *ActivityUseCase) UpdateActivity(req request.UpdateExistingActivity) (res *entities.Activity, httpCode int, errType error, srvError int) {
	//Validate request
	if err := utils.ValidateStruct[request.UpdateExistingActivity](req); err != nil {
		errMsg := fmt.Errorf("%v, reason: %v", interfaceError.InvalidRequestBody, err)
		return nil, fiber.StatusBadRequest, errMsg, utils.HTTPRequestErr
	}

	// Validate activity exist
	if err := u.data.ValidateActivity(req.GetId()); err != nil {
		return nil, fiber.StatusNotFound, err, utils.HTTPRequestErr
	}

	updatedEntities, err := u.data.UpdateActivityData(req)
	if err != nil {
		return nil, fiber.StatusInternalServerError, interfaceError.FailedUpdateExistingActivity, utils.DatabaseError
	}

	return updatedEntities, fiber.StatusOK, nil, 0
}
