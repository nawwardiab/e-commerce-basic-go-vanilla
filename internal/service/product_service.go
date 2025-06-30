package service

import (
	"server/internal/model"
)

// ProductService Interface type (exported)
type ProductService interface {
	Get()([]model.Product, error)
	GetProductByID(id int) (*model.Product, error)
}

// Unexported productRepo interface type
type productRepo interface {
	GetAllProducts()([]model.Product, error)
	GetProductDetails(id int) (*model.Product, error)
}

// Unexported productService that has repo attribute
type productService struct {
    repo productRepo
}

// Constructs new ProductServices object (methods and db access through repo)
func NewProductService(r productRepo) ProductService {
    return &productService{repo: r}
}

// Get – calls repo function that queries db and return slice of all products
func (s productService) Get() ([]model.Product, error) {
    return s.repo.GetAllProducts()
}

// GetProductByID –  calls Repo function that queries db and returns details to one product
func (s productService) GetProductByID(id int) (*model.Product, error){
	return s.repo.GetProductDetails(id)
}