package repository

import (
	"database/sql"
	"fmt"

	"kasir.api/domain"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}

}

func (repo *TransactionRepository) CreateTransaction(items []domain.CheckoutItems) (*domain.Transaction, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	totalAmount := 0
	details := make([]domain.TransactionDetail, 0)

	for _, item := range items {
		var inventoryPrice, stock int
		var inventoryName string

		err := tx.QueryRow("SELECT name, price, stock FROM inventories WHERE id = $1", item.InventoryID).Scan(&inventoryName, &inventoryPrice, &stock)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("inventories id %d not found", item.InventoryID)
		}
		if err != nil {
			return nil, err
		}

		subtotal := inventoryPrice * item.Quantity
		totalAmount += subtotal

		_, err = tx.Exec("UPDATE inventories SET stock = stock - $1 WHERE id = $2", item.Quantity, item.InventoryID)
		if err != nil {
			return nil, err
		}

		details = append(details, domain.TransactionDetail{
			InventoryID:   item.InventoryID,
			InventoryName: inventoryName,
			Quantity:      item.Quantity,
			Subtotal:      subtotal,
		})
	}

	var transactionID int
	err = tx.QueryRow("INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id", totalAmount).Scan(&transactionID)
	if err != nil {
		return nil, err
	}

	for i := range details {
		details[i].TransactionID = transactionID
		_, err = tx.Exec("INSERT INTO transaction_details (transaction_id, inventory_id, quantity, subtotal) VALUES ($1, $2, $3, $4)",
			transactionID, details[i].InventoryID, details[i].Quantity, details[i].Subtotal)
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &domain.Transaction{
		ID:          fmt.Sprintf("%d", transactionID),
		TotalAmount: totalAmount,
		Details:     details,
	}, nil
}
