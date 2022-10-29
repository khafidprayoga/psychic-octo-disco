package tables

import "time"

type Activity struct {
	ID        int        `json:"id" gorm:"size:255;notNull;primaryKey"`
	Email     *string    `json:"email"`
	Title     string     `json:"title"`
	CreatedAt *time.Time `json:"created_at" gorm:"index;autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"index;autoUpdateTime"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"autoUpdateTime"`
}
