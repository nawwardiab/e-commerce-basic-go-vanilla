package repository

import "server/internal/model"


type UserRepo interface {
  GetByUsername(username string) (*model.User, error)
  CreateUser(u *model.User) error
}

type ProductRepo interface {
  GetAllProducts() ([]model.Product, error)
  GetProductDetails(id int) (*model.Product, error)
}