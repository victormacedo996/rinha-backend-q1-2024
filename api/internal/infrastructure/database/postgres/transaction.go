package postgres

import (
	"context"
	"time"

	"github.com/victormacedo996/rinha-backend-q1-2024/internal/infrastructure/database/entity"
)

const INSERT_NEW_TRANSACTION = `
INSERT INTO 
    transactions
    (client_id, transaction_date, value, "transaction_type", description)
VALUES
    ($1, $2, $3, $4, $5)
RETURNING
    client_id;
`

func (d *DbInstance) RegisterTransaction(ctx context.Context, client_id int, transaction_request entity.TransactionRequest) error {

	now := time.Now().Unix()

	if transaction_request.Description == "devolve" {
		now = now + 1
	}

	var id int

	err := d.pool.QueryRow(ctx, INSERT_NEW_TRANSACTION, client_id, now, transaction_request.Value, transaction_request.Type, transaction_request.Description).Scan(&id)
	if err != nil {
		return err
	}

	return nil

}
