// queries database and returns User Object – depends on db connection
package repository

import (
	"fmt"
	"server/internal/model"

	"github.com/jackc/pgx"
)



type UserRepo struct {
  db *pgx.Conn
}

func NewUserRepo(db *pgx.Conn) *UserRepo {
	return &UserRepo{db: db}
}

// GetByUsername uses the db connectio to query users table by username
func (r *UserRepo) GetByUsername(username string) (*model.User, error) {
  u := &model.User{}
  err := r.db.QueryRow(`SELECT * FROM users WHERE username=$1`, username,
  ).Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash)
  if err != nil {
      return nil, fmt.Errorf("GetByUsername: %w", err)
  } else {
    return u, nil
  }
}

// Create uses db connection to query users table and inserts a new user
func (r *UserRepo) CreateUser(u *model.User) error {
  _, err := r.db.Exec(`INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)`,
      u.Username, u.Email, u.PasswordHash,
  )
  if err != nil {
      return fmt.Errorf("create user: %w", err)
  } else {
    return nil
  }
}