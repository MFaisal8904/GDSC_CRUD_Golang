package connection

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// Unit test untuk koneksi database
func TestConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang_database")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}