package product

type ClientGetProductsResponse struct {
	Products []ClientProduct `json:"products"`
	Total    int             `json:"total"`
	Skip     int             `json:"skip"`
	Limit    int             `json:"limit"`
}

type ClientProduct struct {
	ID                   int             `json:"id"`
	Title                string          `json:"title"`
	Description          string          `json:"description"`
	Category             string          `json:"category"`
	Price                float64         `json:"price"`
	DiscountPercentage   float64         `json:"discountPercentage"`
	Rating               float64         `json:"rating"`
	Stock                int             `json:"stock"`
	Tags                 []string        `json:"tags"`
	Brand                string          `json:"brand"`
	SKU                  string          `json:"sku"`
	Weight               int             `json:"weight"`
	Dimensions           ClientDimension `json:"dimensions"`
	WarrantyInformation  string          `json:"warrantyInformation"`
	ShippingInformation  string          `json:"shippingInformation"`
	AvailabilityStatus   string          `json:"availabilityStatus"`
	Reviews              []ClientReview  `json:"reviews"`
	ReturnPolicy         string          `json:"returnPolicy"`
	MinimumOrderQuantity int             `json:"minimumOrderQuantity"`
	Thumbnail            string          `json:"thumbnail"`
	Images               []string        `json:"images"`
}

type ClientDimension struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Depth  float64 `json:"depth"`
}

type ClientReview struct {
	Rating        int    `json:"rating"`
	Comment       string `json:"comment"`
	Date          string `json:"date"`
	ReviewerName  string `json:"reviewerName"`
	ReviewerEmail string `json:"reviewerEmail"`
}
