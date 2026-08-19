package models

import "time"

type Order struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	OrderedAt time.Time `json:"ordered_at" db:"ordered_at"`
}
