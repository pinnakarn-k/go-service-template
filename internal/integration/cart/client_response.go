package cart

import "github.com/shopspring/decimal"

type ClientGetCartsResponse struct {
	Carts []ClientCart `json:"carts"`
	Total int          `json:"total"`
	Skip  int          `json:"skip"`
	Limit int          `json:"limit"`
}

type ClientCart struct {
	ID              int             `json:"id"`
	Products        []ClientProduct `json:"products"`
	Total           decimal.Decimal `json:"total"`
	DiscountedTotal decimal.Decimal `json:"discountedTotal"`
	UserID          int             `json:"userId"`
	TotalProducts   int             `json:"totalProducts"`
	TotalQuantity   int             `json:"totalQuantity"`
}

type ClientProduct struct {
	ID                 int             `json:"id"`
	Title              string          `json:"title"`
	Price              decimal.Decimal `json:"price"`
	Quantity           int             `json:"quantity"`
	Total              decimal.Decimal `json:"total"`
	DiscountPercentage decimal.Decimal `json:"discountPercentage"`
	DiscountedTotal    decimal.Decimal `json:"discountedTotal"`
	Thumbnail          string          `json:"thumbnail"`
}
