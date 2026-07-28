package cart

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

func mapGetCartsResponse(
	src *ClientGetCartsResponse,
) *GetCartsResponse {
	items := make([]CartResponse, 0, len(src.Carts))

	for _, cart := range src.Carts {
		products := make(
			[]CartProductResponse,
			0,
			len(cart.Products),
		)

		for _, product := range cart.Products {
			products = append(products, CartProductResponse{
				ProductID: product.ID,
				Name:      strings.ToUpper(product.Title),
				UnitPrice: formatAmount(product.Price),
				Quantity:  product.Quantity,
				NetAmount: formatAmount(product.DiscountedTotal),
			})
		}

		items = append(items, CartResponse{
			CartID:        cart.ID,
			UserCode:      fmt.Sprintf("USER-%05d", cart.UserID),
			TotalAmount:   formatAmount(cart.Total),
			NetAmount:     formatAmount(cart.DiscountedTotal),
			TotalProducts: cart.TotalProducts,
			TotalQuantity: cart.TotalQuantity,
			Summary: fmt.Sprintf(
				"%d products, %d items",
				cart.TotalProducts,
				cart.TotalQuantity,
			),
			Products: products,
		})
	}

	page := 1
	if src.Limit > 0 {
		page = (src.Skip / src.Limit) + 1
	}

	return &GetCartsResponse{
		Items:      items,
		TotalItems: src.Total,
		Page:       page,
		PageSize:   src.Limit,
	}
}

func formatAmount(amount decimal.Decimal) string {
	return amount.StringFixed(2)
}
