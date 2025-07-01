package handler

import (
	"errors"
	"net/http"
	"server/internal/service"
	"server/internal/session"
	"server/internal/view"
	"strconv"
)

type auth struct {
  userSvc service.UserService
  session session.Session
}

func NewAuthHandler(userService service.UserService, sess *session.Session) *auth{
  return &auth{userSvc: userService, session: *sess}
}

// RegisterHandler – handles http request to create a new user
func (ah *auth) RegisterHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
    _ = view.Render(w, "register.tpl", nil)
    return
  } else if r.Method == http.MethodPost{

    usr := r.FormValue("username")
    email := r.FormValue("email")
    pwd := r.FormValue("password")
    rep := r.FormValue("repeatedPassword")
    if usr == "" || email == "" || pwd == "" || pwd != rep {
      formErr := map[string]string{"Error": "Fill all fields and ensure passwords match"}
      _ = view.Render(w, "register.tpl", formErr)
      return
    }

    _, err := ah.userSvc.Register(usr, email, pwd)
    if err != nil {
      if errors.Is(err, service.ErrUserExist) {
        _ = view.Render(w, "register.tpl", service.ErrUserExist.Error())
        return
      } else {
        http.Error(w, "server error", http.StatusInternalServerError)
        return
      }
    }
    http.Redirect(w, r, "/login", http.StatusSeeOther)
  } else {
    http.Error(w, "server error", http.StatusMethodNotAllowed)
    return 
  }
}

// LoginHandler – handles http request to log in an existing user and redirect to home page.
func (ah *auth) LoginHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
    _ =  view.Render(w, "login.tpl", nil)
    return
  } else if r.Method == http.MethodPost{

    // POST
    username := r.FormValue("username")
    password := r.FormValue("password")
    user, loginErr := ah.userSvc.Login(username, password)
    sessErr := ah.session.Set(w, r, "user_id", strconv.Itoa(user.ID))
    if loginErr != nil {
      code := http.StatusInternalServerError
      if errors.Is(loginErr, service.ErrInvalidCredentials) {
          code = http.StatusUnauthorized
      }
      http.Error(w, loginErr.Error(), code)
      return
    } else if sessErr != nil {
      http.Error(w, "session error", http.StatusInternalServerError)
      return
    } else {
      http.Redirect(w, r, "/", http.StatusSeeOther)
    }
  } else {
    http.Error(w, "server error", http.StatusMethodNotAllowed)
  }
}

// LogoutHandler – removes the session and redirects to login page
func (ah *auth) LogoutHandler(w http.ResponseWriter, r *http.Request){
	err := ah.session.Delete(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	return
	} else {
    http.Redirect(w, r, "/login", http.StatusSeeOther)
  }
}