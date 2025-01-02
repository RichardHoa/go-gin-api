package product

import (
	"net/http"

	"github.com/RichardHoa/go-gin-api/cmd/types"
	"github.com/RichardHoa/go-gin-api/cmd/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) ProductRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.HandleGetProducts).Methods("GET")

}

func (h *Handler) HandleGetProducts(w http.ResponseWriter, r *http.Request) {

	products, err := h.store.GetProducts()
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// for _, product := range products {
	// 	utils.DebuggingPrinting(product)
	// }

	utils.WriteJSONResponse(w, http.StatusOK, products)

}
