package handler

import (
	"net/http"
	"server/internal/view"
)


func (h Handler) HomeHandler(w http.ResponseWriter, r *http.Request){
	// 1. check if user is logged in
	if !h.session.Has(r) {
    http.Redirect(w, r, "/login", http.StatusFound)
    return
  } else {    
    // 2. Prepare template data
    data := map[string]bool{"Logged": true}

    // 3. Render in one step
    if err := view.Render(w, "home.tpl", data); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
  }
}