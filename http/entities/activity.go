package entities

import "time"

type Activity struct {
	ID        int        `json:"id"`
	Email     *string    `json:"email"`
	Title     string     `json:"title"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
