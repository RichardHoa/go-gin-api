package product

import (
	"database/sql"
	"fmt"
	"github.com/RichardHoa/go-gin-api/cmd/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetProducts() ([]types.Product, error) {

	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := make([]types.Product, 0)

	for rows.Next() {
		var product types.Product
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Image,
			&product.Price,
			&product.Quantity,
			&product.CreatedAt,
		); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if len(products) == 0 {
		return nil, fmt.Errorf("no products found")

	}
	return products, nil

}

func (s *Store) CreateProduct(product types.Product) error {

	_, err := s.db.Exec(
		"INSERT INTO products (name, description, image, price, quantity) VALUES (?, ?, ?, ?, ?)",
		product.Name,
		product.Description,
		product.Image,
		product.Price,
		product.Quantity,
	)

	return err

}
