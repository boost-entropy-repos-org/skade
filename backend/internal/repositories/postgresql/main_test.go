package postgresql

import (
	"log"
	"testing"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://skadeuser:skadepw@127.0.0.1:5432/skadedb?sslmode=disable"
)

var testrepo *Repo

func TestMain(m* testing.M) {
	conn, err := sqlx.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Can not connect to DB:", err)
	}

	testrepo = NewRepo(conn)

	os.Exit(m.Run())
}
