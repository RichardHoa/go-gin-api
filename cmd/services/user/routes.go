package user

import (
	"fmt"
	"net/http"

	"github.com/RichardHoa/go-gin-api/cmd/services/auth"
	"github.com/RichardHoa/go-gin-api/cmd/types"
	"github.com/RichardHoa/go-gin-api/cmd/utils"

	// "github.com/go-playground/validator/v10"
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

	err := utils.ParseJSON(r, &payload)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, fmt.Errorf("error parsing JSON: %s", err))
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		friendlyErrors := utils.CreateFriendlyErrorMSG(err)
		utils.WriteJSONResponse(w, http.StatusBadRequest, friendlyErrors)
		return
	}

	_, getUserErr := h.store.GetUserByEmail(payload.Email)
	if getUserErr == nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	hashedPassword, hashPasswordErr := auth.HashPassword(payload.Password)
	if hashPasswordErr != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("error hashing password: %s", hashPasswordErr))
		return
	}

	user := types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	}

	createUserErr := h.store.CreateUser(user)

	utils.DebuggingPrinting(user)

	if createUserErr != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("error creating user: %s", createUserErr))
		return
	}

	sendResErr := utils.WriteJSONResponse(w, http.StatusCreated, map[string]string{
		"message": "New user has been created",
	})

	if sendResErr != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, sendResErr)
		return
	}

}
