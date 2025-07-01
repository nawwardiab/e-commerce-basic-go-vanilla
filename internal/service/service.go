package service

import "server/internal/model"

type UserService interface{
	Register(username, email, password string) (*model.User, error)
	Login(username, password string) (*model.User, error)
}

type ProductService interface {
	Get() ([]model.Product, error)
	GetProductByID(id int) (*model.Product, error)
}