package data

import (
	"github.com/khafidprayoga/psychic-octo-disco/http/entities"
	request "github.com/khafidprayoga/psychic-octo-disco/http/req"
	"github.com/khafidprayoga/psychic-octo-disco/interface/activity/interfaceError"
	"gorm.io/gorm"
	"time"
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
	now := time.Now()
	err = q.dbMysql.Table("activity").Where("id = ?", id).Updates(map[string]any{
		"deleted_at": &now,
	}).Error

	return

}

func (q *ActivityData) DetailActivityData(id string) (data *entities.Activity, err error) {
	if err := q.dbMysql.Where("id = ?", id).Find(&data).Error; err != nil {
		return nil, interfaceError.FailedGetDetailActivity
	}

	return

}

func (q *ActivityData) ListActivityData() (data []entities.Activity, err error) {
	if err := q.dbMysql.Find(&data).Error; err != nil {
		return data, interfaceError.FailedGetListActivity
	}

	return

}
func (q *ActivityData) UpdateActivityData(data request.UpdateExistingTodo) (*entities.Activity, error) {
	return nil, nil

}
func (q *ActivityData) ValidateActivity(id string) (err error) {
	var count int64
	if err := q.dbMysql.Table("activity").
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		return interfaceError.DataNotFound
	}
	return nil

}
