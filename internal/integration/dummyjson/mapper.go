package dummyjson

import (
	"fmt"
	"math"
	"strings"

	"github.com/shopspring/decimal"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// carts
func mapGetCartsResponse(
	src *GetCartsClientResponse,
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

// products
func mapGetProductsResponse(
	src *GetProductsClientResponse,
) *GetProductsResponse {
	if src == nil {
		return nil
	}

	items := make(
		[]ProductResponse,
		0,
		len(src.Products),
	)

	for _, product := range src.Products {
		items = append(
			items,
			ProductResponse{
				ProductID:    product.ID,
				Name:         product.Title,
				CategoryName: mapCategoryName(product.Category),
				Description:  product.Description,
				Price:        formatPrice(product.Price),
				DiscountPrice: calculateDiscountPrice(
					product.Price,
					product.DiscountPercentage,
				),
				StockStatus:  mapStockStatus(product.Stock),
				Rating:       formatRating(product.Rating),
				Brand:        product.Brand,
				ThumbnailURL: product.Thumbnail,
				Tags:         product.Tags,
				Dimensions: DimensionResponse{
					Width:  product.Dimensions.Width,
					Height: product.Dimensions.Height,
					Depth:  product.Dimensions.Depth,
					Volume: calculateVolume(
						product.Dimensions,
					),
				},
				ReviewSummary: mapReviewSummary(
					product.Reviews,
				),
			},
		)
	}

	page := (src.Skip / src.Limit) + 1

	return &GetProductsResponse{
		Items:      items,
		TotalItems: src.Total,
		Page:       page,
		PageSize:   src.Limit,
	}
}

func formatPrice(value float64) string {
	return fmt.Sprintf("%.2f", value)
}

func calculateDiscountPrice(
	price float64,
	discountPercentage float64,
) string {
	discountedPrice := price * (1 - discountPercentage/100)

	return formatPrice(discountedPrice)
}

func mapStockStatus(stock int) string {
	switch {
	case stock <= 0:
		return "OUT_OF_STOCK"

	case stock <= 10:
		return "LOW_STOCK"

	default:
		return "IN_STOCK"
	}
}

func formatRating(rating float64) string {
	return fmt.Sprintf("%.2f/5", rating)
}

func mapCategoryName(category string) string {
	category = strings.ReplaceAll(category, "-", " ")

	return cases.Title(language.English).String(category)
}

func calculateVolume(
	dimensions ProductDimensionsClientResponse,
) float64 {
	return dimensions.Width *
		dimensions.Height *
		dimensions.Depth
}

func mapReviewSummary(
	reviews []ProductReviewClientResponse,
) ReviewSummaryResponse {
	if len(reviews) == 0 {
		return ReviewSummaryResponse{}
	}

	var totalRating int
	latestReview := reviews[0]

	for _, review := range reviews {
		totalRating += review.Rating

		if review.Date > latestReview.Date {
			latestReview = review
		}
	}

	averageRating := float64(totalRating) / float64(len(reviews))

	return ReviewSummaryResponse{
		TotalReviews:   len(reviews),
		AverageRating:  roundToTwoDecimals(averageRating),
		LatestReviewer: latestReview.ReviewerName,
	}
}

func roundToTwoDecimals(value float64) float64 {
	return math.Round(value*100) / 100
}
