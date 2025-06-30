package repository

import (
	"fmt"
	"server/internal/model"
)

// GeAllProducts – queries db products table and returns products details
func (r Repo) GetAllProducts() ([]model.Product, error) {
  rows, err := r.db.Query(`SELECT * FROM products`)
  if err != nil {
    return nil, fmt.Errorf("postgres: ListAllProducts products: %w", err)
  }
  defer rows.Close()

  // slice to hold data from returned rows
  var list []model.Product

  // Loop through rows, using Scan to assign column data to struct fields.
  for rows.Next() {
    var p model.Product
    err := rows.Scan(
      &p.ID,
      &p.Title,
      &p.Year,
      &p.Artist,
      &p.Img,
      &p.Price,
      &p.Genre,
      )
    if err != nil {
      return nil, fmt.Errorf("postgres: scan product: %w", err)
    } else {
      list = append(list, p)
    }
  }
  return list, rows.Err()
}

// GetProductDetails – queries db products table and returns queried product's details
func (r Repo) GetProductDetails(id int) (*model.Product, error) {
  query := `SELECT * FROM products WHERE id=$1`
  row := r.db.QueryRow(query, id)

  var p model.Product
  err := row.Scan(&p.ID, &p.Title, &p.Year, &p.Artist, &p.Img, &p.Price, &p.Genre,)
  if err != nil {
    return nil, fmt.Errorf("postgres: scan product %d: %w", id, err)
  }

  return &p, nil
}