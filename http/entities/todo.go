package entities

import "time"

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
