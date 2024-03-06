package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/config"
	usecase "github.com/victormacedo996/rinha-backend-q1-2024/internal/domain/usecases"
)

type ChiWebserver struct{}

func New() *ChiWebserver {
	return &ChiWebserver{}
}

func setUpRoutes(r *chi.Mux, usecase *usecase.Usecase, validator *validator.Validate) {
	r.Post("/clientes/{id}/transacoes", createTransaction(usecase))
	r.Get("/clientes/{id}/extrato", getBankStatement(usecase))
	r.Get("/health", health)

}

func (cw *ChiWebserver) Start(usecase *usecase.Usecase, validator *validator.Validate) {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	setUpRoutes(router, usecase, validator)

	c := config.GetInstance()
	port := fmt.Sprintf(":%v", c.WebServer.SERVER_PORT)
	fmt.Printf("starting api0%s in port %d", c.WebServer.API_ID, c.WebServer.SERVER_PORT)
	http.ListenAndServe(port, router)
}
