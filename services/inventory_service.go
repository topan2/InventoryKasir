package service

import (
	"errors"

	"kasir.api/domain"
	"kasir.api/repository"
)

type InventoryService struct {
	repo repository.InventoryRepository
}

func NewInventoryService(r repository.InventoryRepository) *InventoryService {
	return &InventoryService{repo: r}
}

func (s *InventoryService) Create(inv *domain.Inventories) error {
	if inv.Stock < 0 {
		return errors.New("stock cannot be negative")
	}
	return s.repo.Create(inv)
}

func (s *InventoryService) GetAll(name string) ([]domain.Inventories, error) {
	return s.repo.GetAll(name)
}

func (s *InventoryService) Update(id string, inv *domain.Inventories) error {
	if inv.Stock < 0 {
		return errors.New("stock cannot be negative")
	}
	return s.repo.Update(id, inv)
}

func (s *InventoryService) Delete(id string) error {
	return s.repo.Delete(id)
}
