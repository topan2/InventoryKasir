package models

type Inventory struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Stock    int      `json:"stock"`
	Category Category `json:"category"`
}
