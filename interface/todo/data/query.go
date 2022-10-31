package data

import (
	"github.com/khafidprayoga/psychic-octo-disco/http/entities"
	"github.com/khafidprayoga/psychic-octo-disco/http/req"
	"github.com/khafidprayoga/psychic-octo-disco/interface/todo/interfaceError"
	"gorm.io/gorm"
	"log"
	"time"
)

type TodoData struct {
	dbMysql gorm.DB
}

func New(db gorm.DB) *TodoData {
	return &TodoData{
		dbMysql: db,
	}
}

func (q *TodoData) CreateNew(req req.CreateNewTodo) (data *entities.Todo, err error) {
	newTodo := &entities.Todo{
		Title:           req.Title,
		ActivityGroupId: *req.ActivityGroupID,
	}

	if err := q.dbMysql.Create(newTodo).Error; err != nil {
		return nil, interfaceError.FailedCreateNewTodo
	}

	return newTodo, nil
}

func (q *TodoData) DeleteTodo(id string) error {
	now := time.Now()
	if err := q.dbMysql.Table("todos").
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"deleted_at": &now,
		}).Error; err != nil {

		log.Println(err)
		return interfaceError.FailedDeleteExistingTodo
	}

	return nil
}

func (q *TodoData) DetailTodo(id string) (*entities.Todo, error) {
	result := new(entities.Todo)
	if err := q.dbMysql.Where("id = ?", id).Find(result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (q *TodoData) ListAllTodo(activityId string) ([]entities.Todo, error) {
	listData := new([]entities.Todo)

	if activityId != "" {
		err := q.dbMysql.Where("activity_group_id = ?", activityId).Find(&listData).Error
		if err != nil {
			return []entities.Todo{}, err
		}
	} else {
		err := q.dbMysql.Find(&listData).Error
		if err != nil {
			return []entities.Todo{}, err
		}
	}

	return *listData, nil
}

func (q *TodoData) UpdateTodo(data req.UpdateExistingTodo) (*entities.Todo, error) {
	var result entities.Todo

	isActive := 0
	newData := entities.Todo{
		Title:    data.Title,
		IsActive: &isActive,
	}

	if data.IsActive != nil {
		isActive := 0
		if *data.IsActive {
			isActive = 1
			newData.IsActive = &isActive
		}

		if !*data.IsActive {
			newData.IsActive = &isActive

		}
	}

	if e := q.dbMysql.Model(&entities.Todo{}).
		Where("id = ?", data.GetId()).
		Updates(newData).Scan(&result).Error; e != nil {
		return nil, e
	}

	return &result, nil
}

func (q *TodoData) ValidateTodo(id string) error {
	var count int64
	if err := q.dbMysql.Table("todos").
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
