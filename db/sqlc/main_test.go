package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/rajatk-31/simplebank/util"
)

var testQueries *Queries
var testDB *sql.DB

// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
// )

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load config", err)
	}
	testDB, err = sql.Open(config.DBdriver, config.DBsource)
	if err != nil {
		log.Fatal("Cannot connect to db")
	}

	testQueries = New(testDB)
	os.Exit(m.Run())

}
