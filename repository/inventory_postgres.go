package repository

import (
	"database/sql"

	"kasir.api/domain"
)

type inventoryPostgres struct {
	db *sql.DB
}

func NewInventoryPostgres(db *sql.DB) InventoryRepository {
	return &inventoryPostgres{db: db}
}

func (r *inventoryPostgres) Create(inv *domain.Inventories) error {
	return r.db.QueryRow(`
        INSERT INTO inventories (name, price, description, stock, category)
        VALUES ($1,$2,$3,$4,$5)
        RETURNING id, created_at, updated_at`,
		inv.Name, inv.Price, inv.Description, inv.Stock, inv.Category,
	).Scan(&inv.ID, &inv.CreatedAt, &inv.UpdatedAt)
}

func (r *inventoryPostgres) GetAll(name string) ([]domain.Inventories, error) {
	rows, err := r.db.Query(`
        SELECT id,name,price,description,stock,category,created_at,updated_at
        FROM inventories WHERE name ILIKE $1`,
		"%"+name+"%",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var inventories []domain.Inventories
	for rows.Next() {
		var inv domain.Inventories
		rows.Scan(
			&inv.ID,
			&inv.Name,
			&inv.Price,
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

func (r *inventoryPostgres) Update(id string, inv *domain.Inventories) error {
	return r.db.QueryRow(`
        UPDATE inventories
        SET name=$1, price=$2, description=$3, stock=$4, category=$5, updated_at=now()
        WHERE id=$6
        RETURNING updated_at`,
		inv.Name, inv.Price, inv.Description, inv.Stock, inv.Category, id,
	).Scan(&inv.UpdatedAt)
}

func (r *inventoryPostgres) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM inventories WHERE id=$1`, id)
	return err
}
