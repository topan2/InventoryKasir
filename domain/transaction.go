package domain

import "time"

type Transaction struct {
	ID          string              `json:"id"`
	TotalAmount int                 `json:"total_amount"`
	Details     []TransactionDetail `json:"details"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
}

type TransactionDetail struct {
	ID            int    `json:"id"`
	TransactionID int    `json:"transaction_id"`
	InventoryID   int    `json:"inventory_id"`
	InventoryName string `json:"inventory_name,omitempty"`
	Type          string `json:"type"` // IN / OUT / SALE
	Quantity      int    `json:"quantity"`
	Subtotal      int    `json:"subtotal"`
}

type CheckoutItems struct {
	InventoryID int `json:"inventory_id"`
	Quantity    int `json:"quantity"`
}

type CheckoutRequest struct {
	Items []CheckoutItems `json:"items"`
}

type DailySalesReport struct {
	Date           string  `json:"date"`
	TotalSales     float64 `json:"total_sales"`
	TotalTransaksi int     `json:"total_transaksi"`
	TotalItem      int     `json:"total_item"`
}
