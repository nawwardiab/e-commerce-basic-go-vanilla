package handler

import (
	"server/internal/service"
	"server/internal/session"
)

// TODO:
// restructure:
// abstract services into a Services struct.
// pass Services type into the handler.
//! update occurences –> e.g. h.userSvc –> h.svcs.userSvc

// Handler struct – abstracts handlers and assign only data each handler requires
type Handler struct{
	svcs service.Services
  session    *session.Session
}

// contructs abstract NewHandler after passing required data
func NewHandler(svcs service.Services, session *session.Session,) *Handler {
	return &Handler{svcs: svcs, session: session,}
}