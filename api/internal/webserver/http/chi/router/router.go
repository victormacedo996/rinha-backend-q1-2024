package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/config"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/infrastructure/database/postgres"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/infrastructure/database/redis"
)

func setUpRoutes(r *chi.Mux, db *postgres.DbInstance, redis *redis.Redis, validator *validator.Validate) {
	r.Use(middleware.Logger)
	r.Post("/clientes/{id}/transacoes", createTransaction(validator, db, redis))
	r.Get("/clientes/{id}/extrato", getBankStatement(db, redis))

	r.Get("/health", health)

}

func Start(db *postgres.DbInstance, redis *redis.Redis, validator *validator.Validate) {
	router := chi.NewRouter()
	setUpRoutes(router, db, redis, validator)

	c := config.GetInstance()
	port := fmt.Sprintf(":%v", c.WebServer.SERVER_PORT)
	http.ListenAndServe(port, router)
}
