package health

import (
	"net/http"

	"github.com/RichardHoa/go-gin-api/cmd/utils"
	"github.com/gorilla/mux"
)

func HealthRoutes(router *mux.Router) {
	router.HandleFunc("/health", HandleHealthCheck).Methods(http.MethodGet)
}

func HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSONResponse(w, http.StatusOK, map[string]string{
		"status": "ok",
	})
}