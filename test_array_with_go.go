package main

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type HTTPClientErrorCode int

type HTTPClientErrorCodeList []HTTPClientErrorCode

type TestArrayWithGo struct {
	SupportedErrorCodes HTTPClientErrorCodeList
}

const (
	insertNewRecord = `insert into test_array_with_go (supported_error_codes)
values($1)
returning supported_error_codes`
)

func NewTestArryWithGo(
	ctx context.Context, trx pgx.Tx,
	supoportedErrorCodes HTTPClientErrorCodeList,
) (TestArrayWithGo, error) {
	row := trx.QueryRow(ctx, insertNewRecord, supoportedErrorCodes)

	var record TestArrayWithGo

	err := row.Scan(&record.SupportedErrorCodes)

	return record, err
}
