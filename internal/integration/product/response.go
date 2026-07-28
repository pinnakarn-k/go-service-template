package product

type GetProductsResponse struct {
	Items      []Product `json:"items"`
	TotalItems int       `json:"totalItems"`
	Page       int       `json:"page"`
	PageSize   int       `json:"pageSize"`
}

type Product struct {
	ProductID     int           `json:"productId"`
	Name          string        `json:"name"`
	CategoryName  string        `json:"categoryName"`
	Description   string        `json:"description"`
	Price         string        `json:"price"`
	DiscountPrice string        `json:"discountPrice"`
	StockStatus   string        `json:"stockStatus"`
	Rating        string        `json:"rating"`
	Brand         string        `json:"brand"`
	ThumbnailURL  string        `json:"thumbnailUrl"`
	Tags          []string      `json:"tags"`
	Dimensions    Dimension     `json:"dimensions"`
	ReviewSummary ReviewSummary `json:"reviewSummary"`
}

type Dimension struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Depth  float64 `json:"depth"`
	Volume float64 `json:"volume"`
}

type ReviewSummary struct {
	TotalReviews   int     `json:"totalReviews"`
	AverageRating  float64 `json:"averageRating"`
	LatestReviewer string  `json:"latestReviewer"`
}
