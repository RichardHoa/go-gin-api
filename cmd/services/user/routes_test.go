package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RichardHoa/go-gin-api/cmd/types"
	"github.com/gorilla/mux"
)

func TestUserHandler(t *testing.T) {

	userStore := &MockUserStore{}
	handler := NewHandler(userStore)

	t.Run("Should fail if user has invalid payload > invalid email", func(t *testing.T) {
		payload := types.UserResgisterPayload{
			FirstName: "Richard",
			LastName:  "Hoa",
			Email:     "invalid email",
			Password:  "2",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(marshalled))

		if err != nil {
			t.Errorf("Error creating request: %s", err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d with msg %s", http.StatusBadRequest, rr.Code, rr.Body.String())
		}
	})

	t.Run("Should fail if user has invalid payload > empty FirstName", func(t *testing.T) {
		payload := types.UserResgisterPayload{
			FirstName: "",
			LastName:  "Hoa",
			Email:     "test@gmail.com",
			Password:  "more than 3 byte and less than 72 bytes",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(marshalled))

		if err != nil {
			t.Errorf("Error creating request: %s", err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d with msg %s", http.StatusBadRequest, rr.Code, rr.Body.String())
		}
	})

	t.Run("Should fail if user has invalid payload > empty LastName", func(t *testing.T) {
		payload := types.UserResgisterPayload{
			FirstName: "SOmething",
			LastName:  "",
			Email:     "test@gmail.com",
			Password:  "more than 3 byte and less than 72 bytes",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(marshalled))

		if err != nil {
			t.Errorf("Error creating request: %s", err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d with msg %s", http.StatusBadRequest, rr.Code, rr.Body.String())
		}
	})

	t.Run("Should fail if user has invalid payload > password less than 3 characters", func(t *testing.T) {
		payload := types.UserResgisterPayload{
			FirstName: "Richard",
			LastName:  "Hoa",
			Email:     "hoa@gmail.com",
			Password:  "a",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(marshalled))

		if err != nil {
			t.Errorf("Error creating request: %s", err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d with msg %s", http.StatusBadRequest, rr.Code, rr.Body.String())
		}
	})

	t.Run("Should fail if user has invalid payload > password more than 72 bytes", func(t *testing.T) {
		payload := types.UserResgisterPayload{
			FirstName: "Richard",
			LastName:  "Hoa",
			Email:     "hoa@gmail.com",
			Password:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit. ehasellus odio. over than 72 bytes",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(marshalled))

		if err != nil {
			t.Errorf("Error creating request: %s", err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d with msg %s", http.StatusBadRequest, rr.Code, rr.Body.String())
		}
	})

	t.Run("Should succeed if user has valid payload", func(t *testing.T) {
		payload := types.UserResgisterPayload{
			FirstName: "Richard",
			LastName:  "Hoa",
			Email:     "hoa@gmail.com",
			Password:  "more than 3 byte and less than 72 bytes",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(marshalled))

		if err != nil {
			t.Errorf("Error creating request: %s", err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d with msg %s", http.StatusCreated, rr.Code, rr.Body.String())
		}
	})
}

type MockUserStore struct {
}

func (s *MockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (s *MockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (s *MockUserStore) CreateUser(user types.User) error {
	return nil
}
