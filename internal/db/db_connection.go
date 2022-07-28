package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var StorageDB *sql.DB

func Init() *sql.DB {

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	databaseName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	connection := fmt.Sprintf(
		"%s:%s@tcp(localhost%s)/%s",
		user, password, port, databaseName,
	)

	StorageDB, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}
	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}
	log.Println("database configured")

	return StorageDB
}
