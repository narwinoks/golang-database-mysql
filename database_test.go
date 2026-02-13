package golang_database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnectionDatabase(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3307)/belajar_golang_database")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func TestName(t *testing.T) {

}
