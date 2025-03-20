package repository

import (
	"database/sql"
	"fmt"

	"github.com/julioCAlmeida/go-api/internal/model"
)

type ProductRepository struct{
	connection *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return ProductRepository{connection: db}
}

func (r *ProductRepository) GetAll() ([]model.Product, error) {
	query := "SELECT id, name, price FROM products ORDER BY id DESC"
	rows, err := r.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var products []model.Product
	var product model.Product

	for rows.Next() {
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}
		products = append(products, product)
	}
	rows.Close()
	return products, nil
}

func (r *ProductRepository) GetByID(id int) (*model.Product, error) {
	query, err := r.connection.Prepare("SELECT id, name, price FROM products WHERE id = $1")
	if err != nil {
		return &model.Product{}, nil
	}

	var product model.Product
	err = query.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	query.Close()
	return &product, nil
}

func (r *ProductRepository) Create(p model.Product) (int, error) {
	var id int
	query, err := r.connection.Prepare("INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id")
	if err != nil {
		return 0, err
	}
	defer query.Close()

	err = query.QueryRow(p.Name, p.Price).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ProductRepository) Update(id int, p model.Product) (*model.Product, error) {
	query, err := r.connection.Prepare(
		"UPDATE products SET name = $2, price = $3 WHERE id = $1 RETURNING id, name, price")
	if err != nil {
		return nil, err
	}
	defer query.Close()

	err = query.QueryRow(p.ID, p.Name, p.Price).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &p, nil
}

func (r *ProductRepository) Delete(id int) error {
	query, err := r.connection.Prepare("DELETE FROM products WHERE id = $1")
	if err != nil {
		return err
	}
	defer query.Close()

	_, err = query.Exec(id)
	if err != nil {		
		return err
	}
	return nil
}