package models

type UpdateInventoryRequest struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       int    `json:"stock"`
	CategoryID  int    `json:"category_id"`
}
