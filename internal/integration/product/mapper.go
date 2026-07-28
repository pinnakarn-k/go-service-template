package product

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func mapGetProductsResponse(
	src *ClientGetProductsResponse,
) *GetProductsResponse {
	if src == nil {
		return nil
	}

	items := make(
		[]Product,
		0,
		len(src.Products),
	)

	for _, product := range src.Products {
		items = append(
			items,
			Product{
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
				Dimensions: Dimension{
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
	dimensions ClientDimension,
) float64 {
	return dimensions.Width *
		dimensions.Height *
		dimensions.Depth
}

func mapReviewSummary(
	reviews []ClientReview,
) ReviewSummary {
	if len(reviews) == 0 {
		return ReviewSummary{}
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

	return ReviewSummary{
		TotalReviews:   len(reviews),
		AverageRating:  roundToTwoDecimals(averageRating),
		LatestReviewer: latestReview.ReviewerName,
	}
}

func roundToTwoDecimals(value float64) float64 {
	return math.Round(value*100) / 100
}
