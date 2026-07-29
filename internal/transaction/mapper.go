package transaction

import (
	"time"

	"github.com/shopspring/decimal"
)

func mapTransactionRecordsToResponses(
	transactionRecords []TransactionRecord,
) []TransactionResponse {
	transactionResponses := make([]TransactionResponse, 0, len(transactionRecords))

	for _, transactionRecord := range transactionRecords {
		transactionResponses = append(transactionResponses, TransactionResponse{
			ID:             transactionRecord.ID,
			AccountNo:      transactionRecord.AccountNo,
			Symbol:         transactionRecord.Symbol,
			Side:           mapSide(transactionRecord.Side),
			Quantity:       formatQuantity(transactionRecord.Quantity),
			Price:          formatAmount(transactionRecord.Price),
			Amount:         formatAmount(transactionRecord.Amount),
			Fee:            formatAmount(transactionRecord.Fee),
			Status:         mapStatus(transactionRecord.Status),
			TradeDate:      formatDate(transactionRecord.TradeDate),
			SettlementDate: formatDate(transactionRecord.SettlementDate),
			CreatedAt:      formatDateTime(transactionRecord.CreatedAt),
			UpdatedAt:      formatDateTime(transactionRecord.UpdatedAt),
		})
	}

	return transactionResponses
}

func mapTransactionRecordToResponse(
	transactionRecord *TransactionRecord,
) *TransactionResponse {
	if transactionRecord == nil {
		return nil
	}

	return &TransactionResponse{
		ID:             transactionRecord.ID,
		AccountNo:      transactionRecord.AccountNo,
		Symbol:         transactionRecord.Symbol,
		Side:           mapSide(transactionRecord.Side),
		Quantity:       formatQuantity(transactionRecord.Quantity),
		Price:          formatAmount(transactionRecord.Price),
		Amount:         formatAmount(transactionRecord.Amount),
		Fee:            formatAmount(transactionRecord.Fee),
		Status:         mapStatus(transactionRecord.Status),
		TradeDate:      formatDate(transactionRecord.TradeDate),
		SettlementDate: formatDate(transactionRecord.SettlementDate),
		CreatedAt:      formatDateTime(transactionRecord.CreatedAt),
		UpdatedAt:      formatDateTime(transactionRecord.UpdatedAt),
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
