package usecase

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/victormacedo996/rinha-backend-q1-2024/internal/domain/entity"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/domain/repository"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/domain/service"
	dblock "github.com/victormacedo996/rinha-backend-q1-2024/internal/infrastructure/dbLock"
)

type Usecase struct {
	repo   *repository.Repository
	dbLock dblock.DbLock
}

var usecase *Usecase

var once sync.Once

func GetInstance(repo *repository.Repository, dblock dblock.DbLock) *Usecase {
	if usecase == nil {
		once.Do(
			func() {
				usecase = &Usecase{
					repo:   repo,
					dbLock: dblock,
				}
			},
		)
	}

	return usecase
}

func (u *Usecase) CreateTransactionUsecase(ctx context.Context, client_id int, incoming_transaction entity.TransactionRequest) (*entity.TransactionResponse, error) {

	err := u.dbLock.LockDb(ctx)
	if err != nil {
		return nil, err
	}
	defer u.dbLock.UnlockDb(ctx)

	client_balance, client_limit, err := u.repo.GetClientBalanceAndLimit(ctx, client_id)
	if err != nil {
		newErr := errors.New("failed to retrieve client balance and limit")
		return nil, errors.Join(newErr, err)
	}
	value := service.CheckDebit(incoming_transaction.Type, incoming_transaction.Value)

	if value+client_balance < client_limit*-1 {
		return nil, errors.New("client exeeds limit")
	}

	err = u.repo.RegisterTransaction(ctx, client_id, incoming_transaction)
	if err != nil {
		newErr := errors.New("error registering transaction")
		return nil, errors.Join(newErr, err)
	}

	client_new_balance := client_balance + value

	err = u.repo.UpdateClientBalance(ctx, client_id, client_new_balance)
	if err != nil {
		newErr := errors.New("error updating client balance")
		return nil, errors.Join(newErr, err)
	}

	transaction_response := &entity.TransactionResponse{
		Limit:   client_limit,
		Balance: client_new_balance,
	}

	return transaction_response, nil

}

func (u *Usecase) GetBankStatement(ctx context.Context, client_id int) (*entity.BankStatement, error) {
	err := u.dbLock.LockDb(ctx)
	if err != nil {
		return nil, err
	}
	defer u.dbLock.UnlockDb(ctx)

	bank_statement, err := u.repo.GetBankStatement(ctx, client_id)
	if err != nil {
		newErr := fmt.Errorf("failed to retrieve client with id %d bank statement", client_id)
		return nil, errors.Join(newErr, err)
	}

	return bank_statement, nil

}
