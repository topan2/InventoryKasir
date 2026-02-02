package repositories

import (
	"database/sql"
	"errors"

	"kasir.api/models"
)

type InventoryRepository struct {
	db *sql.DB
}

func NewInventoryRepository(db *sql.DB) *InventoryRepository {
	return &InventoryRepository{db: db}
}

/*
 * Get all products
 */
func (repo *InventoryRepository) GetAll() ([]models.Inventory, error) {
	query := `
		SELECT
			i.id,
			i.name,
			i.price,
			i.stock,
			c.id,
			c.name,
			c.description
		FROM inventory i
		LEFT JOIN categories c ON i.category = c.id
	`

	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	inventory := make([]models.Inventory, 0)
	//	for rows.Next() {
	//		var i models.Inventory
	//		err := rows.Scan(
	//			&i.ID,
	//			&i.Name,
	//			&i.Price,
	//			&i.Stock,
	//			&i.Category.ID,
	//			&i.Category.Name,
	//			&i.Category.Description,
	//		)

	for rows.Next() {
		var i models.Inventory
		err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Stock,
		)

		if err != nil {
			return nil, err
		}
		inventory = append(inventory, i)
	}

	return inventory, nil
}

/*
 * Create new inventory
 */
func (repo *InventoryRepository) Create(inventory *models.Inventory) error {
	query := `
		INSERT INTO inventory (name, price, stock, category)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	return repo.db.QueryRow(
		query,
		inventory.Name,
		inventory.Stock,
	).Scan(&inventory.ID)
}

/*
 * Get inventory by id
 */
func (repo *InventoryRepository) GetByID(id int) (*models.Inventory, error) {
	query := `
		SELECT
			i.id,
			i.name,
			i.description,
			i.stock,
			c.id,
			c.name,
			c.description
		FROM inventory i
		LEFT JOIN categories c ON i.category = c.id
	WHERE i.id = $1
	`

	var i models.Inventory
	err := repo.db.QueryRow(query, id).Scan(
		&i.ID,
		&i.Name,
		&i.Stock,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("inventory tidak ditemukan")
	}
	if err != nil {
		return nil, err
	}

	return &i, nil
}

/*
 * Update inventory
 */
func (repo *InventoryRepository) Update(inventory *models.Inventory) error {
	query := `
		UPDATE inventory
		SET name = $1,
		    price = $2,
		    stock = $3,
		    category = $4
		WHERE id = $5
	`

	result, err := repo.db.Exec(
		query,
		inventory.Name,
		inventory.Stock,
		inventory.ID,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("inventory tidak ditemukan")
	}

	return nil
}

/*
 * Delete inventory
 */
func (repo *InventoryRepository) Delete(id int) error {
	query := "DELETE FROM inventory WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("inventory tidak ditemukan")
	}

	return err
}
