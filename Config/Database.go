package Config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var (
	DB *sql.DB
)

func InitDB() {
	var (
		DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DATABASE"))
		err error
	)

	DB, err = sql.Open("mysql", DSN)

	if err != nil {
		panic("failed to connect to database with error: " + err.Error())
		return
	}
}
