package database

import (
	"context"

	"github.com/victormacedo996/rinha-backend-q1-2024/internal/infrastructure/database/entity"
)

type Database interface {
	GetClientBalanceAndLimit(context.Context, int) (int, int, error)
	UpdateClientBalance(context.Context, int, int) error
	RegisterTransaction(context.Context, int, entity.TransactionRequest) error
	GetBankStatement(context.Context, int) (*entity.BankStatement, error)
}
