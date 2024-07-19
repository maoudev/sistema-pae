package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log/slog"
)

func connect() *sql.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", "", "", "")

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		slog.Error(err.Error())
	}

	return db
}
