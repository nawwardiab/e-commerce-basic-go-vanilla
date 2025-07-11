package handler

import (
	"net/http"
	"server/internal/middleware"
	"server/internal/model"
	"server/internal/service"
	"server/internal/session"
	"server/internal/view"
	"strconv"

	"gopkg.in/yaml.v3"
)

type CartHandler struct {
  session session.Session
  prodSvc service.ProductService
}

func NewCartHandler(sess *session.Session, prodSvc service.ProductService) *CartHandler{
  return &CartHandler{session: *sess, prodSvc: prodSvc}
}
// CartHandler – loads cart from session, renders products in template
func (c *CartHandler) CartHandler(w http.ResponseWriter, r *http.Request){
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
  middleware.Logger(r)
  
}

// AddToCartHandler – adds an item to cart (prodId and quantity)
func (c *CartHandler) AddToCartHandler(w http.ResponseWriter, r *http.Request){

  if r.Method != http.MethodPost {
    http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    middleware.Logger(r)
    return
    } else {
      // 1. checks if user logged in
      if !c.session.Has(r) {
      http.Redirect(w, r, "/login", http.StatusSeeOther)
      return
    }
    
    pid, _ := strconv.Atoi(r.FormValue("product_id"))
    qty, _ := strconv.Atoi(r.FormValue("quantity"))
    
    // 2. Load existing cart from session
    cart, err := loadCart(&c.session, r)
    if err != nil {
      http.Error(w, "could not load cart", http.StatusInternalServerError)
      return
    }

    // 3. Call pure service
    updatedCart := service.AddToCart(cart, pid, qty)
    
    // 4. Persist updatedCart back into session
    if err := saveCart(&c.session, w, r, updatedCart); err != nil {
      http.Error(w, "cannot save cart", http.StatusBadRequest)
      return
    }
    // 5. Redirect or render
    http.Redirect(w, r, "/cart", http.StatusSeeOther)
    middleware.Logger(r)
  }
}

// RemoveFromCartHandler – removes an item from cart || reduces quantity
func (c *CartHandler) RemoveFromCartHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
      http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
      middleware.Logger(r)
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
    middleware.Logger(r)
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
  loadErr := yaml.Unmarshal([]byte(raw), &cart)
  if loadErr != nil {
      return nil, loadErr
  }
  return cart, nil
}

// saveCart – codes JSON from data and sets 
func saveCart(s *session.Session, w http.ResponseWriter, r *http.Request, cart service.CartMap) error {
    encoded, mashalErr := yaml.Marshal(cart)
    if mashalErr != nil {
      return mashalErr
    }
  return s.Set(w, r, "cart", string(encoded))
}