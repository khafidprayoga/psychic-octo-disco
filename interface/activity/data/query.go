package data

import (
	"github.com/khafidprayoga/psychic-octo-disco/http/entities"
	request "github.com/khafidprayoga/psychic-octo-disco/http/req"
	"github.com/khafidprayoga/psychic-octo-disco/interface/activity/interfaceError"
	"gorm.io/gorm"
)

type ActivityData struct {
	dbMysql gorm.DB
}

func New(db gorm.DB) *ActivityData {
	return &ActivityData{
		dbMysql: db,
	}
}

func (q *ActivityData) CreateNewActivity(req request.CreateNewActivity) (data *entities.Activity, err error) {
	data = &entities.Activity{
		Title: req.Title,
		Email: &req.Email,
	}
	if err := q.dbMysql.Create(data).Error; err != nil {
		return nil, interfaceError.FailedCreateNewActivity
	}

	return
}

func (q *ActivityData) DeleteActivityData(id string) (err error) {
	return

}
func (q *ActivityData) DetailActivityData(id string) (*entities.Activity, error) {
	return nil, nil

}
func (q *ActivityData) ListActivityData(activityId string) ([]entities.Activity, error) {
	return []entities.Activity{}, nil

}
func (q *ActivityData) UpdateActivityData(data request.UpdateExistingTodo) (*entities.Activity, error) {
	return nil, nil

}
func (q *ActivityData) ValidateTodo(id string) (err error) {
	return

}
