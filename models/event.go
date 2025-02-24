package models

import "time"

// swagger:model
type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Location    string    `json:"location" validate:"required"`
	DateTime    time.Time `json:"date_time" validate:"required"`
	UserID      int       `json:"user_id"`
}
