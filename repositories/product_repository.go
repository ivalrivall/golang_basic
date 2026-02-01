package repositories

import (
	"database/sql"
	"errors"
	"kasir-api/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetAll() ([]models.Product, error) {
	query := `
		SELECT p.id, p.name, p.price, p.stock, c.name as category_name, c.description as category_description
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.id
		ORDER BY id ASC
	`
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]models.Product, 0)
	for rows.Next() {
		var p models.Product
		var c models.Category
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &c.Name, &c.Description)
		if err != nil {
			return nil, err
		}
		if c.Name != "" || c.Description != "" {
			p.Category = &c
		}
		products = append(products, p)
	}

	return products, nil
}

func (repo *ProductRepository) Create(product *models.Product) error {
	query := "INSERT INTO products (name, price, stock, category_id) VALUES ($1, $2, $3, $4) RETURNING id"
	if err := repo.db.QueryRow(query, product.Name, product.Price, product.Stock, product.CategoryID).Scan(&product.ID); err != nil {
		return err
	}

	categoryQuery := "SELECT id, name, description FROM categories WHERE id = $1"
	var c models.Category
	if err := repo.db.QueryRow(categoryQuery, product.CategoryID).Scan(&c.ID, &c.Name, &c.Description); err != nil {
		return err
	}
	product.Category = &c
	return nil
}

// GetByID mengambil produk berdasarkan ID.
func (repo *ProductRepository) GetByID(id int) (*models.Product, error) {
	query := `
		SELECT p.id, p.name, p.price, p.stock, p.category_id, c.id, c.name, c.description
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.id
		WHERE p.id = $1
	`

	var p models.Product
	var c models.Category

	err := repo.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID, &c.ID, &c.Name, &c.Description)
	if err == sql.ErrNoRows {
		return nil, errors.New("produk tidak ditemukan")
	}
	if err != nil {
		return nil, err
	}
	if c.Name != "" || c.Description != "" {
		p.Category = &c
	}

	return &p, nil
}

func (repo *ProductRepository) Update(product *models.Product) error {
	query := "UPDATE products SET name = $1, price = $2, stock = $3, category_id = $4 WHERE id = $5"
	result, err := repo.db.Exec(query, product.Name, product.Price, product.Stock, product.CategoryID, product.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("produk tidak ditemukan")
	}

	return nil
}

func (repo *ProductRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("produk tidak ditemukan")
	}

	return err
}
