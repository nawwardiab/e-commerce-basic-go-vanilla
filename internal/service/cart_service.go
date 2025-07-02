package service

import (
	"fmt"
)

// ErrCartService is returned when fetching a cart fails
var ErrCartService = fmt.Errorf("cart not found")


// CartMap representation in session for productId –> quantity
type CartMap map[int]int


// function AddToCart –> returns CartMap (prodID & quantity) and a CartItem object
func AddToCart(cart CartMap, productId, quantity int) CartMap {
	if cart == nil {
		cart = make(CartMap)
	}
	cart[productId] += quantity
	return cart
}

// RemoveFromCart removes the product from cart.
func RemoveFromCart(cart CartMap, productID int) CartMap {
  delete(cart, productID)
  return cart
}


