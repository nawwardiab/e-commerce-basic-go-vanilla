package handler

import "net/http"

type HomeHandler interface {
	HomeHandler(w http.ResponseWriter, r *http.Request)
}

type AuthHandler interface {
	RegisterHandler(w http.ResponseWriter, r *http.Request)
	LoginHandler(w http.ResponseWriter, r *http.Request)
	LogoutHandler(w http.ResponseWriter, r *http.Request)
}

type ProdHandler interface {
	ProductsHandler(w http.ResponseWriter, r *http.Request)
	ProductDetailsHandler(w http.ResponseWriter, r *http.Request)
}

type CartHandler interface {
	CartHandler(w http.ResponseWriter, r *http.Request)
	AddToCartHandler(w http.ResponseWriter, r *http.Request)
	RemoveFromCartHandler(w http.ResponseWriter, r *http.Request)
}
