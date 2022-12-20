package src

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// MySQLに接続する
func NewConnectMySQL(ctx context.Context) (*sql.DB, error) {
	return sql.Open("mysql", C.DatabaseConfig.FormatDSN())
}
