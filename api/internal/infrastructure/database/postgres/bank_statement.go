package postgres

import (
	"context"
	"time"

	dto "github.com/victormacedo996/rinha-backend-q1-2024/internal/dto/response"
)

const SELECT_FIRST_10_TRANSACTIONS_BY_DECREASING_DATE = `
SELECT
    COALESCE (t.transaction_date, 0) AS transaction_date,
    COALESCE (t.value, 0) AS value,
	COALESCE (t.transaction_type, 'd') AS transaction_type,
    COALESCE (t.description, 'warm up') AS description,
    c.balance AS client_balance,
    c.client_limit
FROM
    clients c
LEFT JOIN
    transactions t ON c.id = t.client_id
WHERE
    c.id = $1
ORDER BY
    transaction_date DESC
LIMIT 10;
`

func (d *DbInstance) GetBankStatement(ctx context.Context, client_id int) (*dto.BankStatement, error) {
	rows, err := d.pool.Query(ctx, SELECT_FIRST_10_TRANSACTIONS_BY_DECREASING_DATE, client_id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var latestTransactions []dto.LatestTransactions
	var balance dto.Balance

	for rows.Next() {
		var (
			transaction dto.LatestTransactions
			timestamp   int64
		)

		err = rows.Scan(&timestamp, &transaction.Value, &transaction.Type, &transaction.Description, &balance.Client_balance, &balance.Limit)
		if err != nil {
			return nil, err
		}

		transaction.Carried_out_in = time.Unix(timestamp, 0)

		if transaction.Description != "warm up" {
			latestTransactions = append(latestTransactions, transaction)
		}

	}

	balance.Bank_statement_date = time.Now()

	bank_statement := dto.BankStatement{
		Balance:           balance,
		Last_transactions: latestTransactions,
	}

	return &bank_statement, nil
}
