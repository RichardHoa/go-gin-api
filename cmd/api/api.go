package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/RichardHoa/go-gin-api/cmd/services/product"
	"github.com/RichardHoa/go-gin-api/cmd/services/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	address string
	db      *sql.DB
}

func NewAPIServer(address string, db *sql.DB) *APIServer {
	return &APIServer{
		address: address,
		db:      db,
	}
}

func (server *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(server.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	productStore := product.NewStore(server.db)
	productHandler := product.NewHandler(productStore)
	productHandler.ProductRoutes(subrouter)

	log.Println("Server is running on", server.address)

	return http.ListenAndServe(server.address, router)
}
