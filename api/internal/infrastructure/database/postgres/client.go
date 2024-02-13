package postgres

import "context"

const SELET_USER_BALANCE_AND_LIMIT = `
SELECT 
	balance, 
	client_limit  
FROM 
	clients 
WHERE 
	id = $1;
`

func (d *DbInstance) GetClientBalanceAndLimit(ctx context.Context, client_id int) (int, int, error) {
	var (
		client_balance int
		client_limit   int
	)

	err := d.pool.QueryRow(ctx, SELET_USER_BALANCE_AND_LIMIT, client_id).Scan(&client_balance, &client_limit)
	if err != nil {
		return 0, 0, err
	}

	return client_balance, client_limit, nil
}

const UPDATE_CLIENT_BALANCE = `
UPDATE 
	clients
SET 
	balance = $1
WHERE id = $2
RETURNING
	balance;
`

func (d *DbInstance) UpdateClientBalance(ctx context.Context, client_id int, new_balance int) error {

	var balance int
	err := d.pool.QueryRow(ctx, UPDATE_CLIENT_BALANCE, new_balance, client_id).Scan(&balance)
	if err != nil {
		return err
	}

	return nil
}
