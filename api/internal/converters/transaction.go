package converters

import (
	domainEntity "github.com/victormacedo996/rinha-backend-q1-2024/internal/domain/entity"
	dbEntity "github.com/victormacedo996/rinha-backend-q1-2024/internal/infrastructure/database/entity"
)

func NewTransactionRequestFromDomain(entity domainEntity.TransactionRequest) dbEntity.TransactionRequest {
	return dbEntity.TransactionRequest{
		Value:       entity.Value,
		Type:        entity.Type,
		Description: entity.Description,
	}
}
