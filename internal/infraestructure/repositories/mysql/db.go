package mysql

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {

}

func connect() *sql.DB {
	if err := godotenv.Load(); err != nil {
		println("Error:", err.Error())
	}

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		slog.Error(err.Error())
	}

	return db
}
