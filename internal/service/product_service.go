package service

import (
	"server/internal/model"
	"server/internal/repository"
)

// Unexported productService that has repo attribute
type ProductService struct {
    repo repository.ProductRepo
}

// Constructs new ProductServices object (methods and db access through repo)
func NewProductService(r *repository.ProductRepo) *ProductService {
    return &ProductService{repo: *r}
}

// Get – calls repo function that queries db and return slice of all products
func (s *ProductService) Get() ([]model.Product, error) {
    return s.repo.GetAllProducts()
}

// GetProductByID –  calls Repo function that queries db and returns details to one product
func (s *ProductService) GetProductByID(id int) (*model.Product, error){
	return s.repo.GetProductDetails(id)
}