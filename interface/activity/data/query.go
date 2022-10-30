package data

import (
	"github.com/khafidprayoga/psychic-octo-disco/http/entities"
	"github.com/khafidprayoga/psychic-octo-disco/http/req"
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

func (q *ActivityData) CreateNewActivity(req req.CreateNewTodo) (data *entities.Activity, err error) {
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
func (q *ActivityData) UpdateActivityData(data req.UpdateExistingTodo) (*entities.Activity, error) {
	return nil, nil

}
func (q *ActivityData) ValidateTodo(id string) (err error) {
	return

}
