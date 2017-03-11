package sqlreflect

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
)

// These are functional tests.
// To point these toward a valid database, set $SQLREFLECT_DB in your
// environment.
//
// SQLREFLECT_DB="user=foo dbname=bar" go test ./...

const tCatalog = "sqlreflect"

var dbConnStr = "user=mbutcher dbname=sqlreflect sslmode=disable"
var dbDriverStr = "postgres"
var db *sql.DB

func TestMain(m *testing.M) {
	if len(sql.Drivers()) == 0 {
		fmt.Println("No database drivers for testing")
		os.Exit(1)
	}

	if cstr := os.Getenv("SQLREFLECT_DB"); len(cstr) > 0 {
		dbConnStr = cstr
	}

	c, err := sql.Open(dbDriverStr, dbConnStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := c.Ping(); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	db = c
	if err := setup(c); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	// Run the tests
	exit := m.Run()

	if err := teardown(c); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	c.Close()

	os.Exit(exit)
}

func TestSchemaInfo(t *testing.T) {
	if err := db.Ping(); err != nil {
		t.Error("failed ping")
	}
	opts := NewDBOptions(db, "postgres")
	si := New(opts)
	if !si.Supported() {
		t.Fatal("Unsupported database")
	}
}

func TestSchemaInfo_Tables(t *testing.T) {
	si := New(DBOptions{Driver: "postgres", Queryer: squirrel.NewStmtCacheProxy(db)})
	tables, err := si.Tables("", "")
	if err != nil {
		t.Fatal(err)
	}
	if n := len(tables); n < len(createTables()) {
		t.Errorf("Unexpected number of tables: %d", n)
	}

	// Try again, but with DB name set.
	tables2, err := si.Tables(tCatalog, "")
	if err != nil {
		t.Fatal(err)
	}
	if n := len(tables2); n != len(createTables()) {
		t.Errorf("Unexpected number of tables: %d", n)
		for _, ttt := range tables2 {
			t.Logf("%s.%s", ttt.TableCatalog, ttt.TableNameField)
		}
	}

	found := false
	for _, tt := range tables {
		if tt.TableNameField == "person" {
			found = true
		}
	}
	if !found {
		t.Error("Did not find table 'person'")
	}
}

func TestSchemaInfo_Table(t *testing.T) {
	si := New(DBOptions{Driver: "postgres", Queryer: squirrel.NewStmtCacheProxy(db)})
	table, err := si.Table("person", tCatalog, "")
	if err != nil {
		t.Fatal(err)
	}

	if table.TableNameField != "person" {
		t.Fatal("Expected person, got %q", table.TableNameField)
	}

	// Make sure we can't accidentally look up a view with this command.
	table, err = si.Table("person_name", tCatalog, "")
	if err == nil {
		t.Fatal("Expected to fail table lookup of a view.")
	}
}

func TestSchemaInfo_Views(t *testing.T) {
	si := New(DBOptions{Driver: "postgres", Queryer: squirrel.NewStmtCacheProxy(db)})
	views, err := si.Views("", "")
	if err != nil {
		t.Fatal(err)
	}
	if n := len(views); n != len(createViews()) {
		t.Errorf("Unexpected number of tables: %d", n)
	}

	// TODO: Should probably create a view in the schema creation and test it
	// here.
}

func TestSchemaInfo_View(t *testing.T) {
	si := New(DBOptions{Driver: "postgres", Queryer: squirrel.NewStmtCacheProxy(db)})
	table, err := si.View("person_name", tCatalog, "")
	if err != nil {
		t.Fatal(err)
	}

	if table.TableNameField != "person_name" {
		t.Fatal("Expected person, got %q", table.TableNameField)
	}
}

func setup(db *sql.DB) error {
	for _, s := range createTables() {
		if _, err := db.Exec(s); err != nil {
			fmt.Println("Setup failed. Cleaning up")
			teardown(db)
			return fmt.Errorf("Statement %q failed: %s", s, err)
		}
	}
	for _, s := range createViews() {
		if _, err := db.Exec(s); err != nil {
			fmt.Println("Setup failed. Cleaning up")
			teardown(db)
			return fmt.Errorf("Statement %q failed: %s", s, err)
		}
	}
	return nil
}

func teardown(db *sql.DB) error {
	var last error
	for _, s := range dropViews() {
		if _, err := db.Exec(s); err != nil {
			last = fmt.Errorf("Statement %q failed: %s", s, err)
			fmt.Println(last)
		}
	}
	for _, s := range dropTables() {
		if _, err := db.Exec(s); err != nil {
			last = fmt.Errorf("Statement %q failed: %s", s, err)
			fmt.Println(last)
		}
	}
	return last
}

func createTables() []string {
	return []string{
		`CREATE TABLE person (
			id SERIAL,
			first_name VARCHAR DEFAULT '',
			last_name VARCHAR DEFAULT '',
			PRIMARY KEY (id)
		)`,
		`CREATE TABLE org (
			id SERIAL,
			name VARCHAR DEFAULT '',
			president INTEGER DEFAULT 0 REFERENCES person(id) ON DELETE SET NULL,
			PRIMARY KEY (id)
		)`,
		`CREATE TABLE employees (
			id SERIAL,
			org INTEGER DEFAULT 0 REFERENCES org(id),
			-- Docs suggest this will use primary key. Useful for testing.
			person INTEGER DEFAULT 0 REFERENCES person,
			PRIMARY KEY (id)
		)`,
	}
}

func createViews() []string {
	return []string{
		`CREATE VIEW person_name AS
		SELECT concat(first_name, last_name) AS full_name FROM person;
		`,
	}
}

func dropTables() []string {
	return []string{
		`DROP TABLE employees`,
		`DROP TABLE org`,
		`DROP TABLE person`,
	}
}

func dropViews() []string {
	return []string{
		`DROP VIEW person_name`,
	}
}
