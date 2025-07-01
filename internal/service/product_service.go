package service

import (
	"server/internal/model"
	"server/internal/repository"
)

// Unexported productService that has repo attribute
type productService struct {
    repo repository.ProductRepo
}

// Constructs new ProductServices object (methods and db access through repo)
func NewProductService(r repository.ProductRepo) *productService {
    return &productService{repo: r}
}

// Get – calls repo function that queries db and return slice of all products
func (s *productService) Get() ([]model.Product, error) {
    return s.repo.GetAllProducts()
}

// GetProductByID –  calls Repo function that queries db and returns details to one product
func (s *productService) GetProductByID(id int) (*model.Product, error){
	return s.repo.GetProductDetails(id)
}