package usecase

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/khafidprayoga/psychic-octo-disco/domain"
	"github.com/khafidprayoga/psychic-octo-disco/http/entities"
	request "github.com/khafidprayoga/psychic-octo-disco/http/req"
	"github.com/khafidprayoga/psychic-octo-disco/interface/activity/interfaceError"
	"github.com/khafidprayoga/psychic-octo-disco/utils"
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
		errMsg := fmt.Errorf("%v reason %v", interfaceError.FailedCreateNewActivity, err)
		return nil, fiber.StatusBadRequest, errMsg, utils.HTTPRequestErr
	}

	resData, err := u.data.CreateNewActivity(req)
	if err != nil {
		return nil, fiber.StatusInternalServerError, err, utils.DatabaseError
	}

	return resData, fiber.StatusCreated, nil, 0
}

func (u *ActivityUseCase) DetailActvity(id string) (res *entities.Activity, httpCode int, errType error, srvError int) {
	return
}
func (u *ActivityUseCase) DeleteActivity(id string) (res *entities.Activity, httpCode int, errType error, srvError int) {
	return

}
func (u *ActivityUseCase) ListActivity(id string) (res *entities.Activity, httpCode int, errType error, srvError int) {
	return

}
func (u *ActivityUseCase) UpdateActivity(id string) (res *entities.Activity, httpCode int, errType error, srvError int) {
	return
}
