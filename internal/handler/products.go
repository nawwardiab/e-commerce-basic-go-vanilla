package handler

import (
	"net/http"
	"server/internal/view"
	"strconv"
)

// ProductsHandler – handles http request: fetches all products & renders products page.
func (h Handler) ProductsHandler(w http.ResponseWriter, r *http.Request){
				
	if !h.session.Has(r) {
    http.Redirect(w, r, "/login", http.StatusSeeOther)
    return
  }
 	products, err := h.svcs.ProductSvc.Get()
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
func (h Handler) ProductDetailsHandler(w http.ResponseWriter, r *http.Request){

  if !h.session.Has(r) {
    http.Redirect(w, r, "/login", http.StatusSeeOther)
    return
  }

  id, err := strconv.Atoi(r.PathValue("id"))
  if err != nil {
    http.NotFound(w, r)
    return
  }

  product, err := h.svcs.ProductSvc.GetProductByID(id)
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