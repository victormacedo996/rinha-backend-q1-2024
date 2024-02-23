package main

import (
	"context"
	"fmt"

	"github.com/victormacedo996/rinha-backend-q1-2024/internal/domain/repository"
	usecase "github.com/victormacedo996/rinha-backend-q1-2024/internal/domain/usecases"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/infrastructure/database/postgres"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/infrastructure/dbLock/redis"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/webserver/http/rest/chi/router"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/webserver/http/rest/validator"
)

func main() {

	fmt.Println("Starting server...")
	fmt.Println("")
	fmt.Println("Creating Database instance...")
	db := postgres.GetInstane()
	fmt.Println("Creating Database lock instance...")

	redis := redis.GetInstance()

	fmt.Println("Creating Validator instance...")

	validator := validator.GetInstance()
	ctx := context.Background()

	fmt.Println("Unlocking Database...")
	err := redis.UnlockDb(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("Creating Repository instance...")

	repo := repository.GetInstance(db)
	fmt.Println("Creating Usecase instance...")
	uc := usecase.GetInstance(repo, redis)

	router.New().Start(uc, validator)

}
