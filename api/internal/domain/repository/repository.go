package repository

import (
	"context"
	"sync"

	"github.com/victormacedo996/rinha-backend-q1-2024/internal/converters"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/domain/entity"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/infrastructure/database"
)

type Repository struct {
	db database.Database
}

var repo *Repository

var once sync.Once

func GetInstance(db database.Database) *Repository {
	if repo == nil {
		once.Do(
			func() {
				repo = &Repository{
					db: db,
				}
			},
		)
	}

	return repo
}

func (r *Repository) GetClientBalanceAndLimit(ctx context.Context, client_id int) (int, int, error) {
	client_balance, client_limit, err := r.db.GetClientBalanceAndLimit(ctx, client_id)
	if err != nil {
		return 0, 0, err
	}

	return client_balance, client_limit, nil
}

func (r *Repository) UpdateClientBalance(ctx context.Context, client_id int, new_balance int) error {
	err := r.db.UpdateClientBalance(ctx, client_id, new_balance)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) RegisterTransaction(ctx context.Context, client_id int, transaction_request entity.TransactionRequest) error {
	db_transaction_request := converters.NewTransactionRequestFromDomain(transaction_request)
	err := r.db.RegisterTransaction(ctx, client_id, db_transaction_request)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetBankStatement(ctx context.Context, client_id int) (*entity.BankStatement, error) {
	db_bank_statement, err := r.db.GetBankStatement(ctx, client_id)
	if err != nil {
		return nil, err
	}

	bank_statement := converters.NewBankStatementFromDb(*db_bank_statement)

	return &bank_statement, nil
}
