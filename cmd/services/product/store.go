package product

import (
	"database/sql"
	"fmt"
	"strings"
	"sync"

	"github.com/RichardHoa/go-gin-api/cmd/types"
)

type Store struct {
	db *sql.DB
	mu *sync.Mutex
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
		mu: &sync.Mutex{},
	}
}

func (s *Store) GetProducts() ([]types.Product, error) {

	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []types.Product{}

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

func (s *Store) GetProductsByID(productIDs []int) ([]types.Product, error) {

	placeholders := strings.Repeat(",?", len(productIDs)-1)
	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (?%s)", placeholders)

	// Convert productIDs to []interface{}
	args := make([]interface{}, len(productIDs))
	for i, v := range productIDs {
		args[i] = v
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	products := []types.Product{}

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

func (s *Store) UpdateProduct(product types.Product) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, err := s.db.Exec(
		"UPDATE products SET quantity = ? WHERE id = ?",
		product.Quantity,
		product.ID,
	)

	return err
}