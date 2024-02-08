package postgres

import (
	"context"
	"time"

	dto "github.com/victormacedo996/rinha-backend-q1-2024/internal/webserver/http/chi/dto/response"
)

const GET_FIRST_10_TRANSACTIONS_BY_DATE = `
SELECT
    t.transaction_date,
    t.value,
    t.transaction_type,
    t.description,
    c.balance AS client_balance,
    c.client_limit
FROM
    transactions t
JOIN
    clients c ON t.client_id = c.id
WHERE
    t.client_id = $1
ORDER BY
    t.transaction_date DESC
LIMIT 10;
`

func (d *DbInstance) GetBankStatement(ctx context.Context, client_id int) (*dto.BankStatement, error) {
	rows, err := d.pool.Query(ctx, GET_FIRST_10_TRANSACTIONS_BY_DATE, client_id)
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

		err := rows.Scan(&timestamp, &transaction.Value, &transaction.Type, &transaction.Description, &balance.Client_balance, &balance.Limit)
		if err != nil {
			return nil, err
		}

		transaction.Carried_out_in = time.Unix(timestamp, 0)

		latestTransactions = append(latestTransactions, transaction)
	}

	balance.Bank_statement_date = time.Now()

	bank_statement := dto.BankStatement{
		Balance:           balance,
		Last_transactions: latestTransactions,
	}

	return &bank_statement, nil
}
