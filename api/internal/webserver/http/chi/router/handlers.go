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
	dto "github.com/victormacedo996/rinha-backend-q1-2024/internal/dto/request"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/infrastructure/database/postgres"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/infrastructure/database/redis"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/webserver/http/chi/response"
)

func createTransaction(validator *validator.Validate, db *postgres.DbInstance, redis *redis.Redis) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlId := chi.URLParam(r, "id")
		if urlId == "" {
			response.StatusUnprocessableEntity(w, r, errors.New("Empty id"))
			return
		}

		id, err := strconv.Atoi(urlId)
		if err != nil {
			response.StatusUnprocessableEntity(w, r, errors.New("Invalid id"))
			return
		}

		if id > 5 {
			response.StatusNotFound(w, r, errors.New("User not found"))
			return
		}

		var incoming_transaction dto.TransactionRequest

		err = json.NewDecoder(r.Body).Decode(&incoming_transaction)
		if err != nil {
			response.StatusUnprocessableEntity(w, r, errors.New("Error parsing request body"))
			return
		}

		err = validator.Struct(incoming_transaction)
		if err != nil {
			response.StatusUnprocessableEntity(w, r, errors.New("Error validating request body"))
			return
		}

		data, _ := json.Marshal(incoming_transaction)
		fmt.Println(string(data))

		for {
			lock, err := redis.GetDbLock(r.Context())
			if err != nil {
				response.StatusInternalServerError(w, r, errors.New("Failed to aquire Db lock"))
				return
			}
			if lock != "1" {
				break
			}
		}

		redis.LockDb(r.Context())
		defer redis.UnlockDb(r.Context())

		transaction_response, err := db.RegisterTransaction(r.Context(), id, incoming_transaction)
		if err != nil {
			fmt.Println(err)
			response.StatusUnprocessableEntity(w, r, errors.New("Error creating transaction"))
			return
		}

		response.StatusOk(w, r, transaction_response)
		return
	}
}

func getBankStatement(db *postgres.DbInstance, redis *redis.Redis) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		urlId := chi.URLParam(r, "id")
		if urlId == "" {
			response.StatusUnprocessableEntity(w, r, errors.New("Empty id"))
			return
		}

		id, err := strconv.Atoi(urlId)
		if err != nil {
			response.StatusUnprocessableEntity(w, r, errors.New("Invalid id"))
			return
		}

		if id > 5 {
			response.StatusNotFound(w, r, errors.New("User not found"))
			return
		}

		fmt.Printf("ID: %d", id)

		for {
			lock, err := redis.GetDbLock(r.Context())
			if err != nil {
				response.StatusInternalServerError(w, r, errors.New("Failed to aquire Db lock"))
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
		return
	}

}

func health(w http.ResponseWriter, r *http.Request) {

	c := config.GetInstance()
	response.StatusOk(w, r, map[string]interface{}{"status": "ok", "API_ID": c.WebServer.API_ID})
	return
}
