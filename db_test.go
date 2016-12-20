package sqlreflect

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

// These are functional tests.
// To point these toward a valid database, set $SQLREFLECT_DB in your
// environment.
//
// SQLREFLECT_DB="user=foo dbname=bar" go test ./...

var dbConnStr = "user=mbutcher dbname=sqlreflect sslmode=disable"
var db *sql.DB

func TestMain(m *testing.M) {
	if len(sql.Drivers()) == 0 {
		fmt.Println("No database drivers for testing")
		os.Exit(1)
	}

	if cstr := os.Getenv("SQLREFLECT_DB"); len(cstr) > 0 {
		dbConnStr = cstr
	}

	c, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := c.Ping(); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	db = c
	exit := m.Run()
	c.Close()

	os.Exit(exit)
}

func TestSchemaInfo(t *testing.T) {
	if err := db.Ping(); err != nil {
		t.Error("failed ping")
	}
	t.Skip("not implemented")
}
