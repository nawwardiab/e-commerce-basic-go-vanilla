package handler

import (
	"net/http"
	"server/internal/service"
	"server/internal/session"
	"server/internal/view"
	"strconv"
)

type Prod struct {
  prodSvc service.ProductService
  session session.Session
}


func NewProdHandler(prodSvc service.ProductService, sess *session.Session) *Prod{
  return &Prod{prodSvc: prodSvc, session: *sess}
}

// ProductsHandler – handles http request: fetches all products & renders products page.
func (ph *Prod) ProductsHandler(w http.ResponseWriter, r *http.Request){
				
	if !ph.session.Has(r) {
    http.Redirect(w, r, "/login", http.StatusSeeOther)
    return
  }
 	products, err := ph.prodSvc.Get()
  if err != nil {
    http.Error(w, "server error", http.StatusInternalServerError)
    return
  } else {
    data := map[string]interface{}{"Products": products}
    if err := view.Render(w, "products.tpl", data); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  }
}

// ProductDetailsHandler –  handles http request: fetches requested product details, renders singleProduct.tpl
func (ph *Prod) ProductDetailsHandler(w http.ResponseWriter, r *http.Request){

  if !ph.session.Has(r) {
    http.Redirect(w, r, "/login", http.StatusSeeOther)
    return
  }

  id, err := strconv.Atoi(r.PathValue("id"))
  if err != nil {
    http.NotFound(w, r)
    return
  }

  product, err := ph.prodSvc.GetProductByID(id)
  if err != nil {
    http.Error(w, "server error", http.StatusInternalServerError)
    return
  } else {
    data := map[string]interface{}{"Product": product}
    if err := view.Render(w, "singleProduct.tpl", data); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  }
}