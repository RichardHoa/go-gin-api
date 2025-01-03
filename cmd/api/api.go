package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/RichardHoa/go-gin-api/cmd/services/cart"
	"github.com/RichardHoa/go-gin-api/cmd/services/health"
	"github.com/RichardHoa/go-gin-api/cmd/services/order"
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

	health.HealthRoutes(subrouter)

	userStore := user.NewStore(server.db)
	userHandler := user.NewHandler(userStore)
	userHandler.UserRoutes(subrouter)

	productStore := product.NewStore(server.db)
	productHandler := product.NewHandler(productStore, userStore)
	productHandler.ProductRoutes(subrouter)

	orderStore := order.NewStore(server.db)
	cartHandler := cart.NewHandler(orderStore, productStore, userStore)
	cartHandler.CartRoutes(subrouter)

	log.Println("Server is online at port", server.address)

	return http.ListenAndServe(server.address, router)
}
