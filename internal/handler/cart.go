package handler

import (
	"encoding/json"
	"net/http"
	"server/internal/model"
	"server/internal/service"
	"server/internal/session"
	"server/internal/view"
	"strconv"
)

type cart struct {
  session session.Session
  prodSvc service.ProductService
}

func NewCartHandler(sess *session.Session, prodSvc service.ProductService) *cart{
  return &cart{session: *sess, prodSvc: prodSvc}
}
// CartHandler – loads cart from session, renders products in template
func (c *cart) HandleCart(w http.ResponseWriter, r *http.Request){
  // checks if user logged in
  if !c.session.Has(r) {
    http.Redirect(w, r, "/login", http.StatusSeeOther)
    return
  }
  // load cart from session
  cart, err := loadCart(&c.session, r)
  if err != nil {
      http.Error(w, "cannot load cart", http.StatusSeeOther)
      return
  }

  // fetch each product and append it to CartItem slice
  var items []model.CartItem
  for pid, qty := range cart {
    prod, _ := c.prodSvc.GetProductByID(pid)
    items = append(items, model.CartItem{
      ProductID: pid,
      Quantity:  qty,
      Product:   *prod,
    })
  }

  view.Render(w, "cart.tpl", map[string]interface{}{"CartItems": items})
}

// AddToCartHandler – adds an item to cart (prodId and quantity)
func (c *cart) AddToCartHandler(w http.ResponseWriter, r *http.Request){

  if r.Method != http.MethodPost {
    http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    return
  } else {
    // 1. checks if user logged in
    if !c.session.Has(r) {
      http.Redirect(w, r, "/login", http.StatusSeeOther)
      return
    }

    //! I got id, received product and retracted id from it later (Where is your ear)
    pid, _ := strconv.Atoi(r.FormValue("product_id"))
    qty, _ := strconv.Atoi(r.FormValue("quantity"))

    // 2. Load existing cart from session
    cart, err := loadCart(&c.session, r)
    if err != nil {
      http.Error(w, "could not load cart", http.StatusInternalServerError)
      return
    }

    // 3. Fetch product from productSvc
    prod, err := c.prodSvc.GetProductByID(pid)
    if err != nil {
      http.Error(w, "product not found", http.StatusNotFound)
      return
    }

    // 4. Call pure service
    updatedCart, _ := service.AddToCart(cart, *prod, qty)

    // 5. Persist updatedCart back into session
    if err := saveCart(&c.session, w, r, updatedCart); err != nil {
      http.Error(w, "cannot save cart", http.StatusBadRequest)
      return
    }
    // 6. Redirect or render
    http.Redirect(w, r, "/cart", http.StatusSeeOther)
  }
}

// RemoveFromCartHandler – removes an item from cart || reduces quantity
func (c *cart) RemoveFromCartHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        return
    }
    if !c.session.Has(r) {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    // Parse the product ID to remove
    pid, err := strconv.Atoi(r.FormValue("product_id"))
    if err != nil {
        http.Error(w, "invalid product id", http.StatusBadRequest)
        return
    }

    // Load existing cart
    cart, err := loadCart(&c.session, r)
    if err != nil {
        http.Error(w, "could not load cart", http.StatusInternalServerError)
        return
    }

    // Remove the item
    updatedCart := service.RemoveFromCart(cart, pid)

    // Save the updated cart
    if err := saveCart(&c.session, w, r, updatedCart); err != nil {
        http.Error(w, "could not save cart", http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/cart", http.StatusSeeOther)
}


//! Helpers
// loads existing cartMap (productID → qty) or returns an empty one.
func loadCart(s *session.Session, r *http.Request) (service.CartMap, error) {
  raw, sessErr := s.Get(r, "cart")
  if sessErr != nil {
      return nil, sessErr
  } 
  // if empty, create new
  if raw == "" {
    cMap := make(service.CartMap)
    return cMap, nil
  }

  var cart service.CartMap
  loadErr := json.Unmarshal([]byte(raw), &cart)
  if loadErr != nil {
      return nil, loadErr
  }
  return cart, nil
}

// saveCart – codes JSON from data and sets 
func saveCart(s *session.Session, w http.ResponseWriter, r *http.Request, cart service.CartMap) error {
    encoded, mashalErr := json.Marshal(cart)
    if mashalErr != nil {
      return mashalErr
    }
  return s.Set(w, r, "cart", string(encoded))
}