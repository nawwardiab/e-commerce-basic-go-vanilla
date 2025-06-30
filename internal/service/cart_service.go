package service

import (
	"fmt"
	"server/internal/model"
)

// ErrCartService is returned when fetching a cart fails
var ErrCartService = fmt.Errorf("cart not found")


// CartMap representation in session for productId –> quantity
type CartMap map[int]int


// function AddToCart –> returns CartMap (prodID & quantity) and a CartItem object
func AddToCart(cart CartMap, product model.Product, quantity int) (CartMap, model.CartItem) {
	if cart == nil {
		cart = make(CartMap)
	}
	cart[product.ID] += quantity
	return cart, model.CartItem{
		ProductID: product.ID,
		Quantity: cart[product.ID],
		Product: product,
	}
}

// RemoveFromCart removes the product from cart.
func RemoveFromCart(cart CartMap, productID int) CartMap {
  delete(cart, productID)
  return cart
}

// UpdateQuantity – updates quantity in the CartMap
// func UpdateQuantity(cart CartMap, product model.Product, quantity int)(CartMap, model.CartItem){
// 	if quantity <= 0 {
// 		delete(cart, product.ID)
// 		return cart, model.CartItem{}
// 	}
// 	cart[product.ID] = quantity
// 	return cart, model.CartItem{
// 		ProductID: product.ID,
// 		Quantity: quantity,
// 		Product: product,
// 	}
// }

