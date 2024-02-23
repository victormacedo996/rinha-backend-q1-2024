package entity

import "time"

type Balance struct {
	Client_balance      int
	Bank_statement_date time.Time
	Limit               int
}

type LatestTransactions struct {
	Value          int
	Type           string
	Description    string
	Carried_out_in time.Time
}

type BankStatement struct {
	Balance           Balance
	Last_transactions []LatestTransactions
}
