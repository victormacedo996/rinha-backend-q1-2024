package router

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/infrastructure/database/postgres"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/infrastructure/database/redis"
	dto "github.com/victormacedo996/rinha-backend-q1-2024/internal/webserver/http/chi/dto/request"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/webserver/http/chi/response"
)

func createTransaction(validator *validator.Validate, db *postgres.DbInstance, redis *redis.Redis) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		urlId := chi.URLParam(r, "id")
		if urlId == "" {
			response.StatusBadRequest(w, r, errors.New("Empty id"))
			return
		}

		id, err := strconv.Atoi(urlId)
		if err != nil {
			response.StatusBadRequest(w, r, errors.New("Invalid id"))
			return
		}

		if id > 5 {
			response.StatusNotFound(w, r, errors.New("User not found"))
			return
		}

		var incoming_transaction dto.TransactionRequest

		err = json.NewDecoder(r.Body).Decode(&incoming_transaction)
		if err != nil {
			response.StatusBadRequest(w, r, errors.New("Error parsing request body"))
			return
		}

		err = validator.Struct(incoming_transaction)
		if err != nil {
			response.StatusBadRequest(w, r, errors.New("Error validating request body"))
			return
		}

		for {
			lock := redis.GetDbLock()
			if lock == "" {
				break
			}
		}

		redis.LockDb()

		transaction_response, err := db.RegisterTransaction(r.Context(), id, incoming_transaction)
		if err != nil {
			response.StatusInternalServerError(w, r, errors.New("Error creating transaction"))
			return
		}
		redis.UnlockDb()

		response.StatusOk(w, r, transaction_response)
		return
	}
}

func getBankStatement(w http.ResponseWriter, r *http.Request) {
	response.StatusOk(w, r, map[string]interface{}{"status": "ok"})
	return
}

func health(w http.ResponseWriter, r *http.Request) {
	response.StatusOk(w, r, map[string]interface{}{"status": "ok"})
	return
}
