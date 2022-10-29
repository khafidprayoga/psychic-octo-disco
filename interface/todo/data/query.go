package data

import "gorm.io/gorm"

type TodoData struct {
	dbMysql gorm.DB
}

func New(db gorm.DB) *TodoData {
	return &TodoData{
		dbMysql: db,
	}
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
