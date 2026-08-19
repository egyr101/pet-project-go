package models

type User struct {
	ID           int     `json:"id" db:"id"`
	Name         string  `json:"name" db:"name"`
	Email        string  `json:"email" db:"email"`
	PasswordHash string  `json:"-" db:"password_hash"`
	Balance      float64 `json:"balance" db:"balance"`
}
