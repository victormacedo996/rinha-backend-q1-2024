package postgres

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/victormacedo996/rinha-backend-q1-2024/internal/config"
)

type dbInstance struct {
	pool *pgxpool.Pool
}

var postgres *dbInstance

var once sync.Once

func GetInstane() *dbInstance {
	once.Do(
		func() {
			if postgres == nil {
				config := config.GetInstance()
				conn_string := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?&pool_max_conns=%d&pool_min_conns=%d",
					config.Database.USER,
					config.Database.PWD,
					config.Database.HOST,
					config.Database.PORT,
					config.Database.DATABASE,
					config.Database.MAX_CONS,
					config.Database.MIN_CONS,
				)

				pgconfig, err := pgxpool.ParseConfig(conn_string)

				if err != nil {
					panic(fmt.Errorf("cannot parse db config %s", err))
				}

				dbpool, err := pgxpool.NewWithConfig(context.Background(), pgconfig)
				if err != nil {
					panic(fmt.Errorf("cannot create db pool %s", err))
				}

				postgres = &dbInstance{
					pool: dbpool,
				}
			}
		},
	)

	return postgres
}
