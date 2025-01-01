package user

import (
	"net/http"

	"github.com/RichardHoa/go-gin-api/cmd/types"
	"github.com/RichardHoa/go-gin-api/cmd/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

	var payload types.UserResgisterPayload

	err := utils.ParseJSON(r, payload)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err)
	}

}
