package router

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/infrastructure/postgres"
	dto "github.com/victormacedo996/rinha-backend-q1-2024/internal/webserver/http/chi/dto/request"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/webserver/http/chi/response"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/webserver/http/chi/validator"
)

func createTransaction(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		response.StatusBadRequest(w, r, errors.New("empty id"))
		return
	}

	var incoming_transaction dto.TransactionRequest

	err := json.NewDecoder(r.Body).Decode(&incoming_transaction)
	if err != nil {
		response.StatusBadRequest(w, r, errors.New("Error parsing request body"))
		return
	}

	validator := validator.GetInstance()

	err = validator.Struct(incoming_transaction)
	if err != nil {
		response.StatusBadRequest(w, r, errors.New("Error validating body"))
		return
	}

	db := postgres.GetInstane()

}
