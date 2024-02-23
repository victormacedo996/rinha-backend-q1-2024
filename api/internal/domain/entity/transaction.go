package entity

type TransactionRequest struct {
	Value       int
	Type        string `validate:"eq=c|eq=d"`
	Description string `validate:"min=1,max=10"`
}

type TransactionResponse struct {
	Limit   int
	Balance int
}
