package repository

import "kasir.api/domain"

type InventoryRepository interface {
	Create(inv *domain.Inventories) error
	GetAll(name string) ([]domain.Inventories, error)
	Update(id string, inv *domain.Inventories) error
	Delete(id string) error
}
