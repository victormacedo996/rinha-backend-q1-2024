package converters

import (
	domainEntity "github.com/victormacedo996/rinha-backend-q1-2024/internal/domain/entity"
	dbEntity "github.com/victormacedo996/rinha-backend-q1-2024/internal/infrastructure/database/entity"
)

func NewBankStatementFromDb(db_bank_statement dbEntity.BankStatement) domainEntity.BankStatement {

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
