package db

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewConnectMySQL(ctx context.Context) (*sql.DB, error) {
	return sql.Open("mysql", "")
}
