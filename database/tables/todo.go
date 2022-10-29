package tables

import "time"

type Todo struct {
	ID              int        `json:"id" gorm:"size:255;notNull;primaryKey"`
	ActivityGroupId int        `json:"activity_group_id" gorm:"size:255;notNull"`
	Title           string     `json:"title" gorm:"size:255;notNull"`
	IsActive        *int       `json:"is_active" gorm:"size:255;notNull"`
	Priority        string     `json:"priority" gorm:"size:255;notNull"`
	CreatedAt       *time.Time `json:"created_at" gorm:"index;autoCreateTime"`
	UpdatedAt       *time.Time `json:"updated_at" gorm:"index;autoUpdateTime"`
	DeletedAt       *time.Time `json:"deleted_at" gorm:"null"`
}
