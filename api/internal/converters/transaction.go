package converters

import (
	domainEntity "github.com/victormacedo996/rinha-backend-q1-2024/internal/domain/entity"
	dbEntity "github.com/victormacedo996/rinha-backend-q1-2024/internal/infrastructure/database/entity"
	requestDto "github.com/victormacedo996/rinha-backend-q1-2024/internal/webserver/http/rest/dto/request"
	responseDto "github.com/victormacedo996/rinha-backend-q1-2024/internal/webserver/http/rest/dto/response"
)

func TransactionRequestFromDomainToDb(entity domainEntity.TransactionRequest) dbEntity.TransactionRequest {
	return dbEntity.TransactionRequest{
		Value:       entity.Value,
		Type:        entity.Type,
		Description: entity.Description,
	}
}

func TransactionResponseFromDtoToDomain(dto responseDto.TransactionResponse) domainEntity.TransactionResponse {
	return domainEntity.TransactionResponse{
		Limit:   dto.Limit,
		Balance: dto.Balance,
	}
}

func TransactionRequestFromDtoToDomain(dto requestDto.TransactionRequest) domainEntity.TransactionRequest {
	return domainEntity.TransactionRequest{
		Value:       dto.Value,
		Type:        dto.Type,
		Description: dto.Description,
	}
}

func TransactionResponseFromDomainToDto(entity domainEntity.TransactionResponse) responseDto.TransactionResponse {
	return responseDto.TransactionResponse{
		Limit:   entity.Limit,
		Balance: entity.Balance,
	}
}
