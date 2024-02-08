package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/infrastructure/database/postgres"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/infrastructure/database/redis"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/webserver/http/chi/validator"
)

func setUpRoutes(r chi.Mux) {

	validator := validator.GetInstance()
	db := postgres.GetInstane()
	redis := redis.GetInstance()

	r.Post("/clientes/{id}/transacoes", createTransaction(validator, db, redis))
	r.Get("/clientes/{id}/extrato", getBankStatement)

	r.Get("/health", health)

}
