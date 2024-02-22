package main

import (
	"context"

	"github.com/victormacedo996/rinha-backend-q1-2024/internal/infrastructure/database/postgres"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/infrastructure/dbLock/redis"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/webserver/http/chi/router"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/webserver/http/chi/validator"
)

func main() {

	db := postgres.GetInstane()
	redis := redis.GetInstance()
	validator := validator.GetInstance()
	ctx := context.Background()
	err := redis.UnlockDb(ctx)
	if err != nil {
		panic(err)
	}

	router.Start(db, redis, validator)

}
