// queries database and returns User Object – depends on db connection
package repository

import (
	"fmt"
	"server/internal/model"

	"github.com/jackc/pgx"
)



type userRepo struct {
  db *pgx.Conn
}

func NewUserRepo(db *pgx.Conn) *userRepo {
	return &userRepo{db: db}
}
// TODO: seperate Repos to have each one access only relevant table in db
// GetByUsername uses the db connectio to query users table by username
func (r *userRepo) GetByUsername(username string) (*model.User, error) {
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
func (r *userRepo) CreateUser(u *model.User) error {
  _, err := r.db.Exec(`INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)`,
      u.Username, u.Email, u.PasswordHash,
  )
  if err != nil {
      return fmt.Errorf("create user: %w", err)
  } else {
    return nil
  }
}