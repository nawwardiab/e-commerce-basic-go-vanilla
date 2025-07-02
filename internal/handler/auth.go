package handler

import (
	"errors"
	"net/http"
	"server/internal/middleware"
	"server/internal/service"
	"server/internal/session"
	"server/internal/view"
	"strconv"
)

type AuthHandler struct {
  userSvc service.UserService
  session session.Session
}

func NewAuthHandler(userService service.UserService, sess *session.Session) *AuthHandler{
  return &AuthHandler{userSvc: userService, session: *sess}
}

// RegisterHandler – handles http request to create a new user
func (ah *AuthHandler) GetRegisterHandler(w http.ResponseWriter, r *http.Request){
  middleware.Logger(r)
  _ = view.Render(w, "register.tpl", nil)
    
}

func (ah *AuthHandler) PostRegisterHandler(w http.ResponseWriter, r *http.Request){

  usr := r.FormValue("username")
  email := r.FormValue("email")
  pwd := r.FormValue("password")
  rep := r.FormValue("repeatedPassword")
  if usr == "" || email == "" || pwd == "" || pwd != rep {
    formErr := map[string]string{"Error": "Fill all fields and ensure passwords match"}
    middleware.Logger(r)
    _ = view.Render(w, "register.tpl", formErr)
    return
  }
  
  _, err := ah.userSvc.Register(usr, email, pwd)
  if err != nil {
      if errors.Is(err, service.ErrUserExist) {
      middleware.Logger(r)
      _ = view.Render(w, "register.tpl", service.ErrUserExist.Error())
      return
      } else {
      middleware.Logger(r)
      http.Error(w, "server error", http.StatusInternalServerError)
      return
    }
  }
  middleware.Logger(r)
  http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// LoginHandler – handles http request to log in an existing user and redirect to home page.
func (ah *AuthHandler) GetLoginHandler(w http.ResponseWriter, r *http.Request){
  
    middleware.Logger(r)
    _ =  view.Render(w, "login.tpl", nil)
} 

func (ah *AuthHandler) PostLoginHandler(w http.ResponseWriter, r *http.Request){
  username := r.FormValue("username")
  password := r.FormValue("password")
  user, loginErr := ah.userSvc.Login(username, password)
  sessErr := ah.session.Set(w, r, "user_id", strconv.Itoa(user.ID))
  if loginErr != nil {
    code := http.StatusInternalServerError
    if errors.Is(loginErr, service.ErrInvalidCredentials) {
        code = http.StatusUnauthorized
    }
    middleware.Logger(r)
    http.Error(w, loginErr.Error(), code)
    return
  } else if sessErr != nil {
    middleware.Logger(r)
    http.Error(w, "session error", http.StatusInternalServerError)
    return
  } else {
    middleware.Logger(r)
    http.Redirect(w, r, "/", http.StatusSeeOther)
  }
}

// LogoutHandler – removes the session and redirects to login page
func (ah *AuthHandler) LogoutHandler(w http.ResponseWriter, r *http.Request){
  err := ah.session.Delete(w, r)
	if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    middleware.Logger(r)
    return
    } else {
    middleware.Logger(r)
    http.Redirect(w, r, "/login", http.StatusSeeOther)
  }
}