package db

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx"
)

// ErrDBConnection is returned when the database connection fails.
var ErrDBConnection = errors.New("infra/database: connection failed")

// NewDB initializes and returns a new connection for the given DSN.
func NewDB(dsn string) (*pgx.Conn, error) {

	// parse the URL into a ConnConfig
  config, err := pgx.ParseURI(dsn)
  if err != nil {
    return nil, fmt.Errorf("invalid DSN: %w", err)
  }

	conn, err := pgx.Connect(config)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDBConnection, err)
	}
	return conn, nil
}