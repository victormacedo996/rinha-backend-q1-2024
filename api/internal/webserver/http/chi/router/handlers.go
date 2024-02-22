package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/config"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/domain/service"
	requestDto "github.com/victormacedo996/rinha-backend-q1-2024/internal/dto/request"
	responseDto "github.com/victormacedo996/rinha-backend-q1-2024/internal/dto/response"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/infrastructure/database/postgres"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/infrastructure/dbLock/redis"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/webserver/http/chi/response"
)

func createTransaction(validator *validator.Validate, db *postgres.DbInstance, redis *redis.Redis) http.HandlerFunc {
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

		for {
			lock, err := redis.GetDbLock(r.Context())
			if err != nil {
				response.StatusInternalServerError(w, r, errors.New("failed to aquire Db lock"))
				return
			}
			if lock != "1" {
				break
			}
		}

		redis.LockDb(r.Context())
		defer redis.UnlockDb(r.Context())

		client_balance, client_limit, err := db.GetClientBalanceAndLimit(r.Context(), id)
		if err != nil {
			fmt.Println(err)
			response.StatusUnprocessableEntity(w, r, errors.New("error fetching client balance and limit"))
			return
		}

		value := service.CheckDebit(incoming_transaction.Type, incoming_transaction.Value)

		if value+client_balance < client_limit*-1 {
			response.StatusUnprocessableEntity(w, r, errors.New("client exeeds limit"))
			return
		}

		err = db.RegisterTransaction(r.Context(), id, incoming_transaction)
		if err != nil {
			response.StatusUnprocessableEntity(w, r, errors.New("error registering transaction"))
			return
		}

		client_new_balance := client_balance + value

		err = db.UpdateClientBalance(r.Context(), id, client_new_balance)
		if err != nil {
			response.StatusUnprocessableEntity(w, r, errors.New("error updating client balance"))
			return
		}

		transaction_response := responseDto.TransactionResponse{
			Limit:   client_limit,
			Balance: client_new_balance,
		}

		response.StatusOk(w, r, transaction_response)
	}
}

func getBankStatement(db *postgres.DbInstance, redis *redis.Redis) http.HandlerFunc {

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

		for {
			lock, err := redis.GetDbLock(r.Context())
			if err != nil {
				response.StatusInternalServerError(w, r, errors.New("failed to aquire Db lock"))
				return
			}
			if lock != "1" {
				break
			}
		}

		redis.LockDb(r.Context())
		defer redis.UnlockDb(r.Context())

		bank_statement, err := db.GetBankStatement(r.Context(), id)
		if err != nil {
			fmt.Println(err)
			response.StatusBadRequest(w, r, errors.New("failed to retrieve bank statement"))
			return
		}

		response.StatusOk(w, r, bank_statement)
	}

}

func health(w http.ResponseWriter, r *http.Request) {

	c := config.GetInstance()
	response.StatusOk(w, r, map[string]interface{}{"status": "ok", "API_ID": c.WebServer.API_ID})
}
