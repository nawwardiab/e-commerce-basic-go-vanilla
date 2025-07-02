package service

import (
	"errors"
	"fmt"
	"server/internal/model"
	"server/internal/repository"

	"github.com/jackc/pgx"
	"golang.org/x/crypto/bcrypt"
)

// Errors for handlers to map HTTP code
var ErrInvalidCredentials = errors.New("service: invalid credentials")
var ErrUserExist = errors.New("service: User already exists")


// unexported userService that has repo attribute 
type UserService struct {
  userRepo repository.UserRepo
}

// Creates new UserService object (methods and db access through repo)
func NewUserService(r *repository.UserRepo) *UserService {
  return &UserService{userRepo: *r}
}

// Register – hashes password and returns a new user after querying the database.
func (s *UserService) Register(username, email, password string) (*model.User, error) {
  if existing, _ := s.userRepo.GetByUsername(username); existing != nil {
    return nil, ErrInvalidCredentials
  }

  hashedPwd, err := hashPassword(password) 
  if err != nil {
    return nil, fmt.Errorf("service: hash password: %w", err)
  }
  u := &model.User{
    Username: username,
    Email: email,
    PasswordHash: hashedPwd,
  }
  if err := s.userRepo.CreateUser(u); err != nil {
    return nil, err
  } else {
    return u, nil
  }
}

// Login – queries db and compares input pwd with db pwd.
func (s *UserService) Login(username, password string) (*model.User, error) {
  u, err := s.userRepo.GetByUsername(username)

  if err != nil {
    // only wrap ErrNoRows as invalid creds:
    if errors.Is(err, pgx.ErrNoRows) {
        return nil, ErrInvalidCredentials
    }
    // something worse happened—propagate it
    return nil, fmt.Errorf("service: user lookup: %w", err)
  } 
  if err := checkPassword(u.PasswordHash, password); err != nil {
    return nil, ErrInvalidCredentials
  } else {
    return u, nil
  }
}

// hashPassword – helper function uses bcrypt to hash the password
func hashPassword(password string)(string, error){
  hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  if err != nil {
    return "", err
  } else {
    return string(hashed), nil
  }
}

// checkPassword – helper function to compare db password with input password
func checkPassword(hashed, password string) error{
  return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}