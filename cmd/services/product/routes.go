package product

import (
	"fmt"
	"net/http"

	"github.com/RichardHoa/go-gin-api/cmd/services/auth"
	"github.com/RichardHoa/go-gin-api/cmd/types"
	"github.com/RichardHoa/go-gin-api/cmd/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	ProductStore types.ProductStore
	UserStore    types.UserStore
}

func NewHandler(ProductStore types.ProductStore, UserStore types.UserStore) *Handler {
	return &Handler{
		ProductStore: ProductStore,
		UserStore:    UserStore,
	}
}

func (h *Handler) ProductRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.HandleGetProducts).Methods(http.MethodGet)
	router.HandleFunc("/products", h.HandlePostProducts).Methods(http.MethodPost)

}

func (h *Handler) HandleGetProducts(w http.ResponseWriter, r *http.Request) {

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

	products, err := h.ProductStore.GetProducts()
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, products)

}

func (h *Handler) HandlePostProducts(w http.ResponseWriter, r *http.Request) {

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

	var payload types.ProductCreatePayload

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
	product := types.Product{
		Name:        payload.Name,
		Description: payload.Description,
		Image:       payload.Image,
		Price:       payload.Price,
		Quantity:    payload.Quantity,
	}

	createProductErr := h.ProductStore.CreateProduct(product)
	if createProductErr != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, createProductErr)
		return
	}

	utils.WriteJSONResponse(w, http.StatusCreated, map[string]string{
		"message": fmt.Sprintf("Product %s created successfully", product.Name),
	})

}
