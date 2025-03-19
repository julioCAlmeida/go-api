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

// func (r *ProductRepository) GetByID(id int) (*model.Product, error) {
// 	for _, p := range products {
// 		if p.ID == id {
// 			return &p, nil
// 		}
// 	}
// 	return nil, errors.New("product not found")
// }

// func (r *ProductRepository) Create(p model.Product) model.Product {
// 	products = append(products, p)
// 	return p
// }

// func (r *ProductRepository) Update(id int, p model.Product) (*model.Product, error) {
// 	for i := range products {
// 		if products[i].ID == id {
// 			products[i] = p
// 			return &products[i], nil
// 		}
// 	}
// 	return nil, errors.New("product not found")
// }

// func (r *ProductRepository) Delete(id int) error {
// 	for i, p := range products {
// 		if p.ID == id {
// 			products = append(products[:i], products[i+1:]...)
// 			return nil
// 		}
// 	}
// 	return errors.New("product not found")
// }