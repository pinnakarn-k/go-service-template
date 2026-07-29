package transaction

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"go-service-template/internal/pagination"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) SearchTransactions(
	ctx context.Context,
	request SearchTransactionsRequest,
) (*pagination.Response[TransactionResponse], error) {
	transactionRecords, totalItems, err := s.repository.SearchTransactions(
		ctx,
		request,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"search transactions service: %w",
			err,
		)
	}

	items := mapTransactionRecordsToResponses(transactionRecords)

	return &pagination.Response[TransactionResponse]{
		Data: items,
		Pagination: pagination.NewMeta(
			request.Page,
			request.PageSize,
			totalItems,
		),
	}, nil
}

func (s *Service) GetTransactionByID(
	ctx context.Context,
	id int64,
) (*TransactionResponse, error) {
	transactionRecord, err := s.repository.GetTransactionByID(
		ctx,
		id,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrTransactionNotFound
		}

		return nil, fmt.Errorf(
			"get transaction by id service: %w",
			err,
		)
	}

	return mapTransactionRecordToResponse(transactionRecord), nil
}
