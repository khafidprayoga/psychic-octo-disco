package tables

import "time"

type Todo struct {
	ID              int        `json:"id" gorm:"size:255;notNull;primaryKey"`
	ActivityGroupId int        `json:"activity_group_id"`
	Title           string     `json:"title"`
	IsActive        *string    `json:"is_active"`
	Priority        string     `json:"priority"`
	CreatedAt       *time.Time `json:"created_at" gorm:"index;autoCreateTime"`
	UpdatedAt       *time.Time `json:"updated_at" gorm:"index;autoUpdateTime"`
	DeletedAt       *time.Time `json:"deleted_at" gorm:"autoUpdateTime"`
}
