package transaction

import (
	"time"

	"github.com/shopspring/decimal"
)

func mapTransactionRecordsToResponses(
	src []TransactionRecord,
) []TransactionResponse {
	items := make([]TransactionResponse, 0, len(src))

	for _, transaction := range src {
		items = append(items, TransactionResponse{
			ID:             transaction.ID,
			AccountNo:      transaction.AccountNo,
			Symbol:         transaction.Symbol,
			Side:           mapSide(transaction.Side),
			Quantity:       formatQuantity(transaction.Quantity),
			Price:          formatAmount(transaction.Price),
			Amount:         formatAmount(transaction.Amount),
			Fee:            formatAmount(transaction.Fee),
			Status:         mapStatus(transaction.Status),
			TradeDate:      formatDate(transaction.TradeDate),
			SettlementDate: formatDate(transaction.SettlementDate),
			CreatedAt:      formatDateTime(transaction.CreatedAt),
			UpdatedAt:      formatDateTime(transaction.UpdatedAt),
		})
	}

	return items
}

func mapTransactionRecordToResponse(
	src *TransactionRecord,
) *TransactionResponse {
	if src == nil {
		return nil
	}

	return &TransactionResponse{
		ID:             src.ID,
		AccountNo:      src.AccountNo,
		Symbol:         src.Symbol,
		Side:           mapSide(src.Side),
		Quantity:       formatQuantity(src.Quantity),
		Price:          formatAmount(src.Price),
		Amount:         formatAmount(src.Amount),
		Fee:            formatAmount(src.Fee),
		Status:         mapStatus(src.Status),
		TradeDate:      formatDate(src.TradeDate),
		SettlementDate: formatDate(src.SettlementDate),
		CreatedAt:      formatDateTime(src.CreatedAt),
		UpdatedAt:      formatDateTime(src.UpdatedAt),
	}
}

func mapSide(
	src *string,
) *string {
	if src == nil {
		return nil
	}

	var value string

	switch *src {
	case "B":
		value = "Buy"
	case "S":
		value = "Sell"
	default:
		value = *src
	}

	return &value
}

func mapStatus(
	src *string,
) *string {
	if src == nil {
		return nil
	}

	var value string

	switch *src {
	case "P":
		value = "Pending"
	case "C":
		value = "Completed"
	case "F":
		value = "Failed"
	case "X":
		value = "Cancelled"
	default:
		value = *src
	}

	return &value
}

func formatAmount(
	src *decimal.Decimal,
) *string {
	if src == nil {
		return nil
	}

	value := src.StringFixed(2)

	return &value
}

func formatDate(
	src *time.Time,
) *string {
	if src == nil {
		return nil
	}

	value := src.Format("02/01/2006")

	return &value
}

func formatQuantity(
	src *decimal.Decimal,
) *string {
	if src == nil {
		return nil
	}

	value := src.StringFixed(4)

	return &value
}

func formatDateTime(
	src *time.Time,
) *string {
	if src == nil {
		return nil
	}

	value := src.Format("02/01/2006 15:04:05")

	return &value
}
