package data

import (
	"github.com/khafidprayoga/psychic-octo-disco/http/entities"
	"github.com/khafidprayoga/psychic-octo-disco/http/req"
	"gorm.io/gorm"
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
		ActivityGroupId: req.ActivityGroupID,
	}

	err = q.dbMysql.Create(newTodo).Error
	data = newTodo

	return
}

func (q *TodoData) DeleteTodo(id string) error {
	return nil
}

func (q *TodoData) DetailTodo(id string) error {
	return nil

}

func (q *TodoData) ListAllTodo() error {
	return nil

}

func (q *TodoData) UpdateTodo(id string) error {
	return nil

}

func (q *TodoData) ValidateTodo(id string) error {
	return nil

}
