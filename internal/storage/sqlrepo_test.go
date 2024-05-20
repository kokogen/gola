package storage

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq" // ...
)

// TestDB ...
func TestDB(t *testing.T) {
	t.Helper()
	databaseURL := "host=localhost port=5432 user=postgres password=pg123 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return
}
