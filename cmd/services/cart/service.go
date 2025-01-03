package cart

import (
	"fmt"

	"github.com/RichardHoa/go-gin-api/cmd/types"
)

func GetItemsIDs(items []types.CartItem) ([]int, error) {

	itemIDs := make([]int, len(items))

	for i, item := range items {

		if item.ProductID == 0 {
			return nil, fmt.Errorf("product ID cannot be 0")
		}
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("product quantity must be greater than 0")
		}

		itemIDs[i] = item.ProductID
	}
	return itemIDs, nil

}

func (h *Handler) CreateOrder(products []types.Product, cartItems []types.CartItem, userID int) (
	orderID int,
	totalPrice float64,
	err error,
) {

	productsMap := make(map[int]types.Product, len(products))

	for _, product := range products {
		productsMap[product.ID] = product
	}

	// Check if cart items are in stock
	OutOfStockErr := checkIfCartItemsIsInStock(cartItems, productsMap)
	if OutOfStockErr != nil {
		return 0, 0, OutOfStockErr
	}

	// Calculate total price
	totalPrice = calculateTotalPrice(cartItems, productsMap)

	// Update products quantity
	for _, cartItem := range cartItems {
		DBItem, ok := productsMap[cartItem.ProductID]
		if !ok {
			return 0, 0, fmt.Errorf("product with ID %d not found", cartItem.ProductID)
		}
		DBItem.Quantity -= cartItem.Quantity

		err := h.ProductStore.UpdateProduct(DBItem)
		if err != nil {
			return 0, 0, fmt.Errorf("failed to update product with ID %d: %v", DBItem.ID, err)
		}
	}

	// Create order
	orderID, err = h.OrderStore.CreateOrder(types.Order{
		UserID:  userID,
		Total:   totalPrice,
		Status:  "pending",
		Address: "random address",
	})
	if err != nil {
		return 0, 0, fmt.Errorf("failed to create order: %v", err)
	}

	// Create order items
	for _, item := range cartItems {
		err = h.OrderStore.CreateOrderItem(types.OrderItem{
			OrderID:   orderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     productsMap[item.ProductID].Price,
		})
		if err != nil {
			return 0, 0, fmt.Errorf("failed to create order item: %v", err)
		}
	}

	return orderID, totalPrice, nil
}

func checkIfCartItemsIsInStock(cartItems []types.CartItem, products map[int]types.Product) error {

	if len(cartItems) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, cartItem := range cartItems {
		product, ok := products[cartItem.ProductID]
		if !ok {
			return fmt.Errorf("product with ID %d not found", cartItem.ProductID)
		}

		if product.Quantity < cartItem.Quantity {
			return fmt.Errorf("product with ID %d is out of stock", cartItem.ProductID)
		}

	}

	return nil

}

func calculateTotalPrice(cartItems []types.CartItem, products map[int]types.Product) float64 {

	totalPrice := 0.0

	for _, cartItem := range cartItems {
		product := products[cartItem.ProductID]
		totalPrice += product.Price * float64(cartItem.Quantity)
	}

	return totalPrice

}
