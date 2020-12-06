package postgresql

import (
	"database/sql"
	"log"
	"testing"
	"os"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://skadeuser:test@127.0.0.1:5432/skadedb?sslmode=disable"
)

func TestMain(m* testing.M) {
	_, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Can not connect to DB:", err)
	}

	os.Exit(m.Run())
}
