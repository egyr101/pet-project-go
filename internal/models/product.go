package models

type Product struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Price       int    `json:"price" db:"price"`
	Quantity    int    `json:"quantity" db:"quantity"`
	CategoryID  int    `json:"category_id" db:"category_id"`
}
