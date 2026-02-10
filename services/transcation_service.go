package service

import (
	"kasir.api/domain"
	"kasir.api/repository"
)

type TransactionService struct {
	repo *repository.TransactionRepository
}

func NewTransactionService(r *repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: r}
}

func (s *TransactionService) Checkout(items []domain.CheckoutItems) (*domain.Transaction, error) {
	return s.repo.CreateTransaction(items)
}
