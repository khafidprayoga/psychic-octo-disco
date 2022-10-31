package entities

import (
	"github.com/khafidprayoga/psychic-octo-disco/http/res"
	"time"
)

type Todo struct {
	ID              int        `json:"id"`
	ActivityGroupId int        `json:"activity_group_id"`
	Title           string     `json:"title"`
	IsActive        *int       `json:"is_active" gorm:"default:1"`
	Priority        string     `json:"priority" gorm:"default:very-high"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
}

func (t *Todo) ToResponse() res.CreatedTodo {
	todo := res.CreatedTodo{
		ID:              t.ID,
		ActivityGroupId: t.ActivityGroupId,
		Title:           t.Title,
		IsActive:        false,
		Priority:        t.Priority,
		CreatedAt:       t.CreatedAt,
		UpdatedAt:       t.UpdatedAt,
	}
	if *t.IsActive == 1 {
		todo.IsActive = true
	}
	
	return todo
}
