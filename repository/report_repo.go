package repository

import (
	"database/sql"

	"kasir.api/domain"
)

type ReportRepository struct {
	DB *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{DB: db}
}

func (r *ReportRepository) GetDailySales(date string) (*domain.DailySalesReport, error) {
	var report domain.DailySalesReport

	query := `
	SELECT 
		TO_CHAR(created_at::date, 'YYYY-MM-DD') as date,
		COALESCE(SUM(total),0) as total_sales,
		COUNT(*) as total_transaksi
	FROM transactions
	WHERE created_at::date = $1
	GROUP BY created_at::date
	`

	err := r.DB.QueryRow(query, date).Scan(
		&report.Date,
		&report.TotalSales,
		&report.TotalTransaksi,
	)

	if err != nil {
		// jika tidak ada transaksi
		report.Date = date
		report.TotalSales = 0
		report.TotalTransaksi = 0
	}

	// total item terjual
	queryItem := `
	SELECT COALESCE(SUM(qty),0) 
	FROM transaction_details td
	JOIN transactions t ON td.transaction_id = t.id
	WHERE t.created_at::date = $1
	`
	r.DB.QueryRow(queryItem, date).Scan(&report.TotalItem)

	return &report, nil
}

func (r *ReportRepository) GetSalesReport(startDate, endDate string) (domain.SalesReport, error) {
	var report domain.SalesReport
	report.StartDate = startDate
	report.EndDate = endDate

	// total revenue + total transaksi
	query := `
		SELECT 
			COALESCE(SUM(total_amount),0) as total_revenue,
			COUNT(*) as total_transaksi
		FROM transactions
		WHERE LOWER(type)='sale'
		AND created_at::date BETWEEN $1 AND $2
	`

	err := r.DB.QueryRow(query, startDate, endDate).Scan(&report.TotalRevenue, &report.TotalTransaksi)
	if err != nil {
		return report, err
	}

	// best seller
	bestQuery := `
	SELECT 
        i.id,
        i.name,
        COALESCE(SUM(td.quantity),0) AS qty_terjual
    FROM transaction_details td
    JOIN inventories i ON i.id = td.inventories_id
    JOIN transactions t ON t.id = td.transaction_id
    WHERE t.created_at::date BETWEEN $1 AND $2
    GROUP BY i.id, i.name
    ORDER BY qty_terjual DESC
    LIMIT 1;
	`

	var best domain.BestSeller
	err = r.DB.QueryRow(bestQuery, startDate, endDate).Scan(&best.InventoryID, &best.Name, &best.TotalQty)
	if err == sql.ErrNoRows {
		report.ProdukTerlaris = domain.BestSeller{
			Name:     "-",
			TotalQty: 0,
		}
		return report, nil
	}
	if err != nil {
		return report, err
	}

	report.ProdukTerlaris = domain.BestSeller{
		InventoryID: best.InventoryID,
		Name:        best.Name,
		TotalQty:    best.TotalQty,
	}

	return report, nil
}
