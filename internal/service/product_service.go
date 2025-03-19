package service

import (
	"github.com/julioCAlmeida/go-api/internal/model"
	"github.com/julioCAlmeida/go-api/internal/repository"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(r repository.ProductRepository) ProductService {
	return ProductService{repo: r}
}

func (s *ProductService) GetAll() ([]model.Product, error) {
	return s.repo.GetAll()
}

// func (s *ProductService) GetByID(id int) (*model.Product, error) {
// 	return s.repo.GetByID(id)
// }

// func (s *ProductService) Create(p model.Product) model.Product {
// 	return s.repo.Create(p)
// }

// func (s *ProductService) Update(id int, p model.Product) (*model.Product, error) {
// 	return s.repo.Update(id, p)
// }

// func (s *ProductService) Delete(id int) error {
// 	return s.repo.Delete(id)
// }