package domain

type BestSeller struct {
	InventoryID int    `json:"inventory_id"`
	Name        string `json:"name"`
	TotalQty    int    `json:"total_qty"`
}

type SalesReport struct {
	StartDate      string     `json:"start_date"`
	EndDate        string     `json:"end_date"`
	TotalRevenue   float64    `json:"total_revenue"`
	TotalTransaksi int        `json:"total_transaksi"`
	ProdukTerlaris BestSeller `json:"best_seller"`
}
