package dto

import "time"

type LatestTransactions struct {
	Value          int       `json:"valor"`
	Type           string    `json:"tipo"`
	Description    string    `json:"descricao"`
	Carried_out_in time.Time `json:"realizada_em"`
}

type BankStatement struct {
	Client_balance      int                  `json:"total"`
	Bank_statement_date time.Time            `json:"data_extrato"`
	Limit               int                  `json:"limite"`
	Last_transactions   []LatestTransactions `json:"ultimas_transacoes"`
}
