package transaction

import (
	"time"

	"github.com/shopspring/decimal"
)

type TransactionRecord struct {
	ID             *int64
	AccountNo      *string
	Symbol         *string
	Side           *string
	Quantity       *decimal.Decimal
	Price          *decimal.Decimal
	Amount         *decimal.Decimal
	Fee            *decimal.Decimal
	Status         *string
	TradeDate      *time.Time
	SettlementDate *time.Time
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}
