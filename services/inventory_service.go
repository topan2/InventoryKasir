package services

import (
	"kasir.api/models"
	"kasir.api/repositories"
)

type InventoryService struct {
	repo *repositories.InventoryRepository
}

func NewInventoryService(repo *repositories.InventoryRepository) *InventoryService {
	return &InventoryService{repo: repo}
}

func (s *InventoryService) GetAll() ([]models.Inventory, error) {
	return s.repo.GetAll()
}

func (s *InventoryService) Create(req models.CreateInventoryRequest) (*models.Inventory, error) {
	inventory := &models.Inventory{
		Name:  req.Name,
		Stock: req.Stock,
		Category: models.Category{
			ID: req.CategoryID,
		},
	}

	if err := s.repo.Create(inventory); err != nil {
		return nil, err
	}

	return inventory, nil
}

func (s *InventoryService) GetByID(id int) (*models.Inventory, error) {
	return s.repo.GetByID(id)
}

func (s *InventoryService) Update(id int, req models.UpdateInventoryRequest) (*models.Inventory, error) {
	inventory := &models.Inventory{
		ID:    id,
		Name:  req.Name,
		Stock: req.Stock,
		Category: models.Category{
			ID: req.CategoryID,
		},
	}

	if err := s.repo.Update(inventory); err != nil {
		return nil, err
	}

	return inventory, nil
}

func (s *InventoryService) Delete(id int) error {
	return s.repo.Delete(id)
}
