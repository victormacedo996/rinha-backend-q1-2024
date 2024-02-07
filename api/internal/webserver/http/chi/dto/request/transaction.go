package dto

type TransactionRequest struct {
	Value       int    `json:"valor"`
	Type        string `json:"tipo" validate:"eq=c|eq=d"`
	Description string `json:"descricao" validate:"max=10"`
}
