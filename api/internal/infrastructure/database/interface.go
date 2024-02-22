package database

import (
	"context"

	dtoRequest "github.com/victormacedo996/rinha-backend-q1-2024/internal/dto/request"
	dtoResponse "github.com/victormacedo996/rinha-backend-q1-2024/internal/dto/response"
)

type Database interface {
	GetClientBalanceAndLimit(context.Context, int) (int, int, error)
	UpdateClientBalance(context.Context, int, int) error
	RegisterTransaction(context.Context, int, dtoRequest.TransactionRequest) error
	GetBankStatement(context.Context, int) (*dtoResponse.BankStatement, error)
}
