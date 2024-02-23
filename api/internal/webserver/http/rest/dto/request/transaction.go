package dto

type TransactionRequest struct {
	Value       int    `json:"valor"`
	Type        string `json:"tipo" validate:"eq=c|eq=d"`
	Description string `json:"descricao" validate:"min=1,max=10"`
}
