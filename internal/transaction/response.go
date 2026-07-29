package transaction

type TransactionResponse struct {
	ID             *int64  `json:"id"`
	AccountNo      *string `json:"accountNo"`
	Symbol         *string `json:"symbol"`
	Side           *string `json:"side"`
	Quantity       *string `json:"quantity"`
	Price          *string `json:"price"`
	Amount         *string `json:"amount"`
	Fee            *string `json:"fee"`
	Status         *string `json:"status"`
	TradeDate      *string `json:"tradeDate"`
	SettlementDate *string `json:"settlementDate"`
	CreatedAt      *string `json:"createdAt"`
	UpdatedAt      *string `json:"updatedAt"`
}
