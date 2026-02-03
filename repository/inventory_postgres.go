package repository

import (
	"database/sql"

	"kasir.api/domain"
)

type inventoryPostgres struct {
	db *sql.DB
}

func NewInventoryPostgres(db *sql.DB) InventoryRepository {
	return &inventoryPostgres{db}
}

func (r *inventoryPostgres) Create(inv *domain.Inventory) error {
	return r.db.QueryRow(`
        INSERT INTO inventories (name, description, stock, category)
        VALUES ($1,$2,$3,$4)
        RETURNING id, created_at, updated_at`,
		inv.Name, inv.Description, inv.Stock, inv.Category,
	).Scan(&inv.ID, &inv.CreatedAt, &inv.UpdatedAt)
}

func (r *inventoryPostgres) FindAll() ([]domain.Inventory, error) {
	rows, err := r.db.Query(`
        SELECT id,name,description,stock,category,created_at,updated_at
        FROM inventories`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var inventories []domain.Inventory
	for rows.Next() {
		var inv domain.Inventory
		rows.Scan(
			&inv.ID,
			&inv.Name,
			&inv.Description,
			&inv.Stock,
			&inv.Category,
			&inv.CreatedAt,
			&inv.UpdatedAt,
		)
		inventories = append(inventories, inv)
	}
	return inventories, nil
}

func (r *inventoryPostgres) Update(id string, inv *domain.Inventory) error {
	return r.db.QueryRow(`
        UPDATE inventories
        SET name=$1, description=$2, stock=$3, category=$4, updated_at=now()
        WHERE id=$5
        RETURNING updated_at`,
		inv.Name, inv.Description, inv.Stock, inv.Category, id,
	).Scan(&inv.UpdatedAt)
}

func (r *inventoryPostgres) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM inventories WHERE id=$1`, id)
	return err
}
