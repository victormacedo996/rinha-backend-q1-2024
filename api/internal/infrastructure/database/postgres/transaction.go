package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	dtoRequest "github.com/victormacedo996/rinha-backend-q1-2024/internal/webserver/http/chi/dto/request"
	dtoResponse "github.com/victormacedo996/rinha-backend-q1-2024/internal/webserver/http/chi/dto/response"
)

const INSERT_NEW_TRANSACTION = `
WITH new_transaction AS (

    INSERT INTO transactions (client_id, transaction_date, value, transaction_type, description)
    SELECT
        @client_id::integer AS client_id,
        @transaction_date AS transaction_date, 
        @value AS value, 
        @transaction_type AS transaction_type,
        @description AS description
    WHERE 
        (SELECT balance FROM clients WHERE id = @client_id) + @value >= -(SELECT client_limit FROM clients WHERE id = @client_id)
        
    RETURNING client_id, value
    )
    UPDATE clients
        SET balance = balance + new_transaction.value
        FROM new_transaction
        WHERE clients.id = new_transaction.client_id
        RETURNING clients.balance AS new_balance, clients.client_limit;
`

func (d *DbInstance) RegisterTransaction(ctx context.Context, client_id int, transaction_request dtoRequest.TransactionRequest) (*dtoResponse.TransactionResponse, error) {

	now := time.Now().Unix()

	var transaction_response dtoResponse.TransactionResponse

	err := d.pool.QueryRow(ctx, INSERT_NEW_TRANSACTION, pgx.NamedArgs{"client_id": client_id, "transaction_date": now, "value": transaction_request.Value, "transaction_type": transaction_request.Type, "description": transaction_request.Description}).Scan(&transaction_response.Balance, &transaction_response.Limit)
	if err != nil {
		return nil, err
	}

	return &transaction_response, nil

}
