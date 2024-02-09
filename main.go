package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	db  *pgxpool.Pool
	ctx = context.Background()
)

func initDB() error {
	var (
		err  error
		conf *pgxpool.Config
	)

	conf, err = pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}

	conf.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		dataTypeNames := []string{
			"http_client_error_code",
			"_http_client_error_code",
			"http_client_error_code_list",
			"_http_client_error_code_list",
		}

		for _, typeName := range dataTypeNames {
			dataType, err := conn.LoadType(ctx, typeName)
			if err != nil {
				return err
			}
			conn.TypeMap().RegisterType(dataType)
		}

		return nil
	}

	db, err = pgxpool.NewWithConfig(context.Background(), conf)

	return err
}

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	trx, err := db.Begin(ctx)
	if err != nil {
		panic(err)
	}

	rec, err := NewTestArryWithGo(ctx, trx, HTTPClientErrorCodeList{400})
	if err != nil {
		_ = trx.Rollback(ctx)

		panic(err)
	}

	_ = trx.Commit(ctx)

	fmt.Printf("%#v\n", rec)
}
