package router

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/config"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/converters"
	usecase "github.com/victormacedo996/rinha-backend-q1-2024/internal/domain/usecases"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/webserver/http/rest/chi/response"
	requestDto "github.com/victormacedo996/rinha-backend-q1-2024/internal/webserver/http/rest/dto/request"
)

func createTransaction(validator *validator.Validate, usecase *usecase.Usecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlId := chi.URLParam(r, "id")
		if urlId == "" {
			response.StatusUnprocessableEntity(w, r, errors.New("empty id"))
			return
		}

		id, err := strconv.Atoi(urlId)
		if err != nil {
			response.StatusUnprocessableEntity(w, r, errors.New("invalid id"))
			return
		}

		if id > 5 {
			response.StatusNotFound(w, r, errors.New("user not found"))
			return
		}

		var incoming_transaction requestDto.TransactionRequest

		err = json.NewDecoder(r.Body).Decode(&incoming_transaction)
		if err != nil {
			response.StatusUnprocessableEntity(w, r, errors.New("error parsing request body"))
			return
		}

		err = validator.Struct(incoming_transaction)
		if err != nil {
			response.StatusUnprocessableEntity(w, r, errors.New("error validating request body"))
			return
		}

		domain_transaction := converters.TransactionRequestFromDtoToDomain(incoming_transaction)

		created_transaction, err := usecase.CreateTransactionUsecase(r.Context(), id, domain_transaction)
		if err != nil {
			response.StatusUnprocessableEntity(w, r, err)
			return
		}

		transaction_response := converters.TransactionResponseFromDomainToDto(*created_transaction)

		response.StatusOk(w, r, transaction_response)
	}
}

func getBankStatement(usecase *usecase.Usecase) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		urlId := chi.URLParam(r, "id")
		if urlId == "" {
			response.StatusUnprocessableEntity(w, r, errors.New("empty id"))
			return
		}

		id, err := strconv.Atoi(urlId)
		if err != nil {
			response.StatusUnprocessableEntity(w, r, errors.New("invalid id"))
			return
		}

		if id > 5 {
			response.StatusNotFound(w, r, errors.New("user not found"))
			return

		}

		bank_statement, err := usecase.GetBankStatement(r.Context(), id)
		if err != nil {
			response.StatusBadRequest(w, r, err)
			return
		}

		response_bank_statement := converters.BankStatementFromDomainToDto(*bank_statement)

		response.StatusOk(w, r, response_bank_statement)
	}

}

func health(w http.ResponseWriter, r *http.Request) {

	c := config.GetInstance()
	response.StatusOk(w, r, map[string]interface{}{"status": "ok", "API_ID": c.WebServer.API_ID})
}
