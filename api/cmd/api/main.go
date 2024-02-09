package main

import (
	"context"
	"fmt"

	"github.com/victormacedo996/rinha-backend-q1-2024/internal/config"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/infrastructure/database/postgres"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/infrastructure/database/redis"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/webserver/http/chi/router"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/webserver/http/chi/validator"
)

func main() {

	db := postgres.GetInstane()
	redis := redis.GetInstance()
	validator := validator.GetInstance()
	c := config.GetInstance()
	ctx := context.Background()
	fmt.Println("unlocking db")
	err := redis.UnlockDb(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("db unlocked")

	router.Start(db, redis, validator)
	fmt.Println(fmt.Sprintf("started api0%s", c.WebServer.API_ID))

}
