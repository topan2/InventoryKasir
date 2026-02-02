package models

type CreateInventoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
	CategoryID  int    `json:"category_id" binding:"required"`
}
