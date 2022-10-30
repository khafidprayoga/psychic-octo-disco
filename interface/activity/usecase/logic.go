package usecase

import (
	"github.com/khafidprayoga/psychic-octo-disco/domain"
	"github.com/khafidprayoga/psychic-octo-disco/http/entities"
)

type ActivityUseCase struct {
	data domain.ActivityData
}

func New(data domain.ActivityData) *ActivityUseCase {
	return &ActivityUseCase{
		data,
	}
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
