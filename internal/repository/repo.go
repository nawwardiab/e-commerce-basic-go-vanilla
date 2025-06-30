package repository

import "github.com/jackc/pgx"

type Repo struct {
  db *pgx.Conn
}

func NewRepo(db *pgx.Conn) *Repo {
  return &Repo{db: db}
}