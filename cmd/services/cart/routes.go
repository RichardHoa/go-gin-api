package cart

import (
	"fmt"
	"net/http"

	"github.com/RichardHoa/go-gin-api/cmd/services/auth"
	"github.com/RichardHoa/go-gin-api/cmd/types"
	"github.com/RichardHoa/go-gin-api/cmd/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	OrderStore   types.OrderStore
	ProductStore types.ProductStore
	UserStore    types.UserStore
}

func NewHandler(OrderStore types.OrderStore, ProductStore types.ProductStore, UserStore types.UserStore) *Handler {
	return &Handler{
		OrderStore:   OrderStore,
		ProductStore: ProductStore,
		UserStore:    UserStore,
	}
}

func (h *Handler) CartRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", h.HandleCheckOut).Methods(http.MethodPost)

}

func (h *Handler) HandleCheckOut(w http.ResponseWriter, r *http.Request) {

	// Authenticate user has valid token
	userID, getUserTokenErr := auth.AuthenticateUserToken(r)
	if getUserTokenErr != nil {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, getUserTokenErr)
		return
	}

	_, getUserErr := h.UserStore.GetUserByID(userID)

	if getUserErr != nil {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, getUserErr)
		return
	}

	var payload types.CartCheckOutPayload

	err := utils.ParseJSON(r, &payload)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, fmt.Errorf("JSON format error: %s", err))
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		friendlyErrors := utils.CreateFriendlyErrorMSG(err)
		utils.WriteJSONResponse(w, http.StatusBadRequest, friendlyErrors)
		return
	}

	// Get product items IDs
	itemIDs, err := GetItemsIDs(payload.Items)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// Get products
	products, err := h.ProductStore.GetProductsByID(itemIDs)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Create order
	orderId, totalPrice, err := h.CreateOrder(products, payload.Items, userID)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	// Return order ID
	utils.WriteJSONResponse(w, http.StatusOK, map[string]any{
		"orderId":    orderId,
		"totalPrice": totalPrice,
	})

}
