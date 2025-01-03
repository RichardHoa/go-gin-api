package product

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RichardHoa/go-gin-api/cmd/services/auth"
	"github.com/RichardHoa/go-gin-api/cmd/types"
	"github.com/gorilla/mux"
)

func TestProductRoutes(t *testing.T) {

	userStore := &MockUserStore{}
	productStore := &MockProductStore{}
	handler := NewHandler(productStore, userStore)
	secret := []byte("secret")

	t.Run("Get product with valid token > ok", func(t *testing.T) {

		req, err := http.NewRequest("GET", "/products", bytes.NewBuffer(nil))
		if err != nil {
			t.Errorf("Error creating request: %s", err)
		}

		token, _ := auth.GenerateJWT(secret, 1)
		req.Header.Set("Authorization", "Bearer "+token)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/products", handler.HandleGetProducts)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d with msg %s", http.StatusOK, rr.Code, rr.Body.String())
		}
	})

	t.Run("Get product with invalid signature token > unauthorized", func(t *testing.T) {

		req, err := http.NewRequest("GET", "/products", bytes.NewBuffer(nil))
		if err != nil {
			t.Errorf("Error creating request: %s", err)
		}

		token, _ := auth.GenerateJWT([]byte("invalid secret"), 1)
		req.Header.Set("Authorization", "Bearer "+token)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/products", handler.HandleGetProducts)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized {
			t.Errorf("Expected status code %d, got %d with msg %s", http.StatusUnauthorized, rr.Code, rr.Body.String())
		}
	})

	t.Run("Get product with missing token > unauthorized", func(t *testing.T) {

		req, err := http.NewRequest("GET", "/products", bytes.NewBuffer(nil))
		if err != nil {
			t.Errorf("Error creating request: %s", err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/products", handler.HandleGetProducts)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusUnauthorized {
			t.Errorf("Expected status code %d, got %d with msg %s", http.StatusUnauthorized, rr.Code, rr.Body.String())
		}
	})

	t.Run("Post product with valid token > created", func(t *testing.T) {

		payload := types.ProductCreatePayload{
			Name:        "Product 1",
			Description: "Description 1",
			Price:       10.0,
			Quantity:    10,
			Image:       "https://example.com/image.jpg",
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest("POST", "/products", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Errorf("Error creating request: %s", err)
		}

		token, _ := auth.GenerateJWT(secret, 1)
		req.Header.Set("Authorization", "Bearer "+token)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/products", handler.HandlePostProducts)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d with msg %s", http.StatusCreated, rr.Code, rr.Body.String())
		}
	})

	t.Run("Post product with missting Name payload > bad request", func(t *testing.T) {

		payload := types.ProductCreatePayload{
			Description: "Description 1",
			Price:       10.0,
			Quantity:    10,
			Image:       "https://example.com/image.jpg",
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest("POST", "/products", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Errorf("Error creating request: %s", err)
		}

		token, _ := auth.GenerateJWT(secret, 1)
		req.Header.Set("Authorization", "Bearer "+token)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/products", handler.HandlePostProducts)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d with msg %s", http.StatusBadRequest, rr.Code, rr.Body.String())
		}
	})

	t.Run("Post product with price == 0 > bad request", func(t *testing.T) {

		payload := types.ProductCreatePayload{
			Name:        "Product 1",
			Description: "Description 1",
			Price:       0,
			Quantity:    10,
			Image:       "https://example.com/image.jpg",
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest("POST", "/products", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Errorf("Error creating request: %s", err)
		}

		token, _ := auth.GenerateJWT(secret, 1)
		req.Header.Set("Authorization", "Bearer "+token)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/products", handler.HandlePostProducts)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d with msg %s", http.StatusBadRequest, rr.Code, rr.Body.String())
		}
	})

	t.Run("Post product with quantity == 0 > bad request", func(t *testing.T) {

		payload := types.ProductCreatePayload{
			Name:        "Product 1",
			Description: "Description 1",
			Price:       10.00,
			Quantity:    0,
			Image:       "https://example.com/image.jpg",
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest("POST", "/products", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Errorf("Error creating request: %s", err)
		}

		token, _ := auth.GenerateJWT(secret, 1)
		req.Header.Set("Authorization", "Bearer "+token)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/products", handler.HandlePostProducts)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d with msg %s", http.StatusBadRequest, rr.Code, rr.Body.String())
		}
	})

}

type MockUserStore struct{}

func (s *MockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (s *MockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (s *MockUserStore) CreateUser(user types.User) error {
	return nil
}

type MockProductStore struct{}

func (s *MockProductStore) GetProducts() ([]types.Product, error) {
	return nil, nil
}

func (s *MockProductStore) CreateProduct(product types.Product) error {
	return nil
}

func (s *MockProductStore) GetProductsByID(productIDs []int) ([]types.Product, error) {
	return nil, nil
}

func (s *MockProductStore) UpdateProduct(product types.Product) error {
	return nil
}
