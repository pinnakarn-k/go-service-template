package dummyjson

// cart
type GetCartsResponse struct {
	Items      []CartResponse `json:"items"`
	TotalItems int            `json:"totalItems"`
	Page       int            `json:"page"`
	PageSize   int            `json:"pageSize"`
}

type CartResponse struct {
	CartID        int                   `json:"cartId"`
	UserCode      string                `json:"userCode"`
	TotalAmount   string                `json:"totalAmount"`
	NetAmount     string                `json:"netAmount"`
	TotalProducts int                   `json:"totalProducts"`
	TotalQuantity int                   `json:"totalQuantity"`
	Summary       string                `json:"summary"`
	Products      []CartProductResponse `json:"products"`
}

type CartProductResponse struct {
	ProductID int    `json:"productId"`
	Name      string `json:"name"`
	UnitPrice string `json:"unitPrice"`
	Quantity  int    `json:"quantity"`
	NetAmount string `json:"netAmount"`
}
