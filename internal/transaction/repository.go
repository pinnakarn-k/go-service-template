package transaction

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"go-service-template/internal/pagination"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) SearchTransactions(
	ctx context.Context,
	request SearchTransactionsRequest,
) ([]TransactionRecord, int, error) {
	const operation = "search transactions repository"

	const countQuery = `
		SELECT COUNT(*)
		FROM transactions
		WHERE
			(NULLIF($1, '') IS NULL OR account_no = $1)
			AND (
				NULLIF($2, '') IS NULL
				OR trade_date >= NULLIF($2, '')::date
			)
			AND (
				NULLIF($3, '') IS NULL
				OR trade_date < NULLIF($3, '')::date + INTERVAL '1 day'
			)
	`

	var totalItems int

	err := r.db.QueryRowContext(
		ctx,
		countQuery,
		request.AccountNo,
		request.FromDate,
		request.ToDate,
	).Scan(&totalItems)
	if err != nil {
		return nil, 0, fmt.Errorf(
			"%s [count]: %w",
			operation,
			err,
		)
	}

	const searchQuery = `
		SELECT
			id,
			account_no,
			symbol,
			side,
			quantity,
			price,
			amount,
			fee,
			status,
			trade_date,
			settlement_date,
			created_at,
			updated_at
		FROM transactions
		WHERE
			(NULLIF($1, '') IS NULL OR account_no = $1)
			AND (
				NULLIF($2, '') IS NULL
				OR trade_date >= NULLIF($2, '')::date
			)
			AND (
				NULLIF($3, '') IS NULL
				OR trade_date < NULLIF($3, '')::date + INTERVAL '1 day'
			)
		ORDER BY trade_date DESC, id DESC
		LIMIT $4
		OFFSET $5
	`

	offset := pagination.Offset(
		request.Page,
		request.PageSize,
	)

	rows, err := r.db.QueryContext(
		ctx,
		searchQuery,
		request.AccountNo,
		request.FromDate,
		request.ToDate,
		request.PageSize,
		offset,
	)
	if err != nil {
		return nil, 0, fmt.Errorf(
			"%s [search]: %w",
			operation,
			err,
		)
	}
	defer rows.Close()

	transactionRecords := make([]TransactionRecord, 0)

	for rows.Next() {
		var transactionRecord TransactionRecord

		err := rows.Scan(
			&transactionRecord.ID,
			&transactionRecord.AccountNo,
			&transactionRecord.Symbol,
			&transactionRecord.Side,
			&transactionRecord.Quantity,
			&transactionRecord.Price,
			&transactionRecord.Amount,
			&transactionRecord.Fee,
			&transactionRecord.Status,
			&transactionRecord.TradeDate,
			&transactionRecord.SettlementDate,
			&transactionRecord.CreatedAt,
			&transactionRecord.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf(
				"%s [scan]: %w",
				operation,
				err,
			)
		}

		transactionRecords = append(
			transactionRecords,
			transactionRecord,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf(
			"%s [iterate]: %w",
			operation,
			err,
		)
	}

	return transactionRecords, totalItems, nil
}

func (r *Repository) GetTransactionByID(
	ctx context.Context,
	id int64,
) (*TransactionRecord, error) {
	const query = `
		SELECT
			id,
			account_no,
			symbol,
			side,
			quantity,
			price,
			amount,
			fee,
			status,
			trade_date,
			settlement_date,
			created_at,
			updated_at
		FROM transactions
		WHERE id = $1
	`

	var transactionRecord TransactionRecord

	err := r.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&transactionRecord.ID,
		&transactionRecord.AccountNo,
		&transactionRecord.Symbol,
		&transactionRecord.Side,
		&transactionRecord.Quantity,
		&transactionRecord.Price,
		&transactionRecord.Amount,
		&transactionRecord.Fee,
		&transactionRecord.Status,
		&transactionRecord.TradeDate,
		&transactionRecord.SettlementDate,
		&transactionRecord.CreatedAt,
		&transactionRecord.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}

		return nil, fmt.Errorf(
			"get transaction by id repository [scan]: %w",
			err,
		)
	}

	return &transactionRecord, nil
}
