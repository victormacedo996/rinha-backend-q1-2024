package webserver

import (
	"github.com/go-playground/validator/v10"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/infrastructure/database/postgres"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/infrastructure/dbLock/redis"
)

type Webserver interface {
	Start(*postgres.DbInstance, *redis.Redis, *validator.Validate)
}
