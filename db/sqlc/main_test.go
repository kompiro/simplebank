package db

import (
	"context"
	"log"
	"testing"

	"github.com/jackc/pgx/v5"
)

const (
	dbSource = "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := pgx.Connect(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: " + err.Error())
	}

	testQueries = New(conn)

	m.Run()
}
