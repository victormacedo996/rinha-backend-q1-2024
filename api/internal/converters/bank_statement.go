package converters

import (
	domainEntity "github.com/victormacedo996/rinha-backend-q1-2024/internal/domain/entity"
	dbEntity "github.com/victormacedo996/rinha-backend-q1-2024/internal/infrastructure/database/entity"
	responseDto "github.com/victormacedo996/rinha-backend-q1-2024/internal/webserver/http/rest/dto/response"
)

func BankStatementFromDbToDomain(db_bank_statement dbEntity.BankStatement) domainEntity.BankStatement {

	balance := domainEntity.Balance{
		Client_balance:      db_bank_statement.Balance.Client_balance,
		Bank_statement_date: db_bank_statement.Balance.Bank_statement_date,
		Limit:               db_bank_statement.Balance.Limit,
	}

	var latest_transactions []domainEntity.LatestTransactions

	for _, value := range db_bank_statement.Last_transactions {
		transaction := domainEntity.LatestTransactions{
			Value:          value.Value,
			Type:           value.Type,
			Description:    value.Description,
			Carried_out_in: value.Carried_out_in,
		}

		latest_transactions = append(latest_transactions, transaction)
	}

	return domainEntity.BankStatement{
		Balance:           balance,
		Last_transactions: latest_transactions,
	}
}

func BankStatementFromDomainToDto(domain_bank_statement domainEntity.BankStatement) responseDto.BankStatement {
	balance := responseDto.Balance{
		Client_balance:      domain_bank_statement.Balance.Client_balance,
		Bank_statement_date: domain_bank_statement.Balance.Bank_statement_date,
		Limit:               domain_bank_statement.Balance.Limit,
	}

	var latest_transactions []responseDto.LatestTransactions

	for _, value := range domain_bank_statement.Last_transactions {
		transaction := responseDto.LatestTransactions{
			Value:          value.Value,
			Type:           value.Type,
			Description:    value.Description,
			Carried_out_in: value.Carried_out_in,
		}

		latest_transactions = append(latest_transactions, transaction)
	}

	return responseDto.BankStatement{
		Balance:           balance,
		Last_transactions: latest_transactions,
	}
}
