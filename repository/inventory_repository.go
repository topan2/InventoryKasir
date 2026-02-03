package repository

import "kasir.api/domain"

type InventoryRepository interface {
	Create(inv *domain.Inventory) error
	FindAll() ([]domain.Inventory, error)
	Update(id string, inv *domain.Inventory) error
	Delete(id string) error
}
