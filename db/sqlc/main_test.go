package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)
var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://skadeuser:test@127.0.0.1:5432/skadedb?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to Database:", err)
	}
	testQueries = New(conn)

	os.Exit(m.Run())
}