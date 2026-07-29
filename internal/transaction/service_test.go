package transaction

import (
	"context"
	"database/sql"
	"errors"
	"testing"
)

type mockTransactionRepository struct {
	searchTransactionsFunc func(
		ctx context.Context,
		request SearchTransactionsRequest,
	) ([]TransactionRecord, int, error)

	getTransactionByIDFunc func(
		ctx context.Context,
		id int64,
	) (*TransactionRecord, error)
}

var _ transactionRepository = (*mockTransactionRepository)(nil)

func (m *mockTransactionRepository) SearchTransactions(
	ctx context.Context,
	request SearchTransactionsRequest,
) ([]TransactionRecord, int, error) {
	return m.searchTransactionsFunc(ctx, request)
}

func (m *mockTransactionRepository) GetTransactionByID(
	ctx context.Context,
	id int64,
) (*TransactionRecord, error) {
	return m.getTransactionByIDFunc(ctx, id)
}

func TestService_GetTransactionByID_Success(t *testing.T) {
	t.Parallel()

	expectedID := int64(1)
	expectedAccountNo := "ACC-001"

	repository := &mockTransactionRepository{
		getTransactionByIDFunc: func(
			ctx context.Context,
			id int64,
		) (*TransactionRecord, error) {
			if id != expectedID {
				t.Fatalf(
					"expected id %d, got %d",
					expectedID,
					id,
				)
			}

			return &TransactionRecord{
				ID:        &expectedID,
				AccountNo: &expectedAccountNo,
			}, nil
		},
	}

	service := NewService(repository)

	transactionResponse, err := service.GetTransactionByID(
		context.Background(),
		expectedID,
	)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if transactionResponse == nil {
		t.Fatal("expected transaction response, got nil")
	}

	if transactionResponse.ID == nil {
		t.Fatal("expected transaction id, got nil")
	}

	if *transactionResponse.ID != expectedID {
		t.Errorf(
			"expected transaction id %d, got %d",
			expectedID,
			*transactionResponse.ID,
		)
	}

	if transactionResponse.AccountNo == nil {
		t.Fatal("expected account no, got nil")
	}

	if *transactionResponse.AccountNo != expectedAccountNo {
		t.Errorf(
			"expected account no %q, got %q",
			expectedAccountNo,
			*transactionResponse.AccountNo,
		)
	}
}

func TestService_GetTransactionByID_NotFound(t *testing.T) {
	t.Parallel()

	expectedID := int64(999)

	repository := &mockTransactionRepository{
		getTransactionByIDFunc: func(
			ctx context.Context,
			id int64,
		) (*TransactionRecord, error) {
			if id != expectedID {
				t.Fatalf(
					"expected id %d, got %d",
					expectedID,
					id,
				)
			}

			return nil, sql.ErrNoRows
		},
	}

	service := NewService(repository)

	transactionResponse, err := service.GetTransactionByID(
		context.Background(),
		expectedID,
	)

	if transactionResponse != nil {
		t.Fatal("expected nil transaction response")
	}

	if !errors.Is(err, ErrTransactionNotFound) {
		t.Fatalf(
			"expected ErrTransactionNotFound, got %v",
			err,
		)
	}
}

func TestService_GetTransactionByID_RepositoryError(t *testing.T) {
	t.Parallel()

	expectedID := int64(1)

	repositoryErr := errors.New("database unavailable")

	repository := &mockTransactionRepository{
		getTransactionByIDFunc: func(
			ctx context.Context,
			id int64,
		) (*TransactionRecord, error) {
			if id != expectedID {
				t.Fatalf(
					"expected id %d, got %d",
					expectedID,
					id,
				)
			}

			return nil, repositoryErr
		},
	}

	service := NewService(repository)

	transactionResponse, err := service.GetTransactionByID(
		context.Background(),
		expectedID,
	)

	if transactionResponse != nil {
		t.Fatal("expected nil transaction response")
	}

	if err == nil {
		t.Fatal("expected error")
	}

	if !errors.Is(err, repositoryErr) {
		t.Fatalf(
			"expected wrapped repository error, got %v",
			err,
		)
	}
}
