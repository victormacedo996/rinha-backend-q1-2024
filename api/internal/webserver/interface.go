package webserver

import (
	"github.com/go-playground/validator/v10"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/infrastructure/database"
	dblock "github.com/victormacedo996/rinha-backend-q1-2024/internal/infrastructure/dbLock"
)

type Webserver interface {
	Start(*database.Database, *dblock.DbLock, *validator.Validate)
}
