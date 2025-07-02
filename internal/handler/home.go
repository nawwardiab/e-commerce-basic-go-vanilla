package handler

import (
	"net/http"
	"server/internal/middleware"
	"server/internal/session"
	"server/internal/view"
)

type HomeHandler struct {
  session session.Session
}

func NewHomeHandler(sess *session.Session) *HomeHandler{
  return &HomeHandler{session: *sess}
}

func (h *HomeHandler) HomeHandler(w http.ResponseWriter, r *http.Request){
	// 1. check if user is logged in
	if !h.session.Has(r) {
    http.Redirect(w, r, "/login", http.StatusFound)
    middleware.Logger(r)
    return
  } else {    
    // 2. Prepare template data
    data := map[string]bool{"Logged": true}

    // 3. Render in one step
    if err := view.Render(w, "home.tpl", data); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    middleware.Logger(r)

  }
}