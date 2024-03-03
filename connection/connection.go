package connection

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// koneksi ke database
func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang_database")
	if err != nil {
		panic(err)
	}

	return db

}
