package main

import (
	"context"
	"log"
	"os"

	logadapter "github.com/jackc/pgx-zerolog"
	"github.com/rs/zerolog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"

	"github.com/techschool/simplebank/api"
	db "github.com/techschool/simplebank/db/sqlc"
)

const (
	dbSource      = "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:3000"
)

func main() {
	config, err := pgxpool.ParseConfig(dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: " + err.Error())
	}

	zlogger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()

	adapter := logadapter.NewLogger(zlogger)

	tracer := tracelog.TraceLog{
		Logger:   adapter,
		LogLevel: tracelog.LogLevelTrace,
	}
	config.ConnConfig.Tracer = &tracer

	connPool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal("cannot connect to db: " + err.Error())
	}

	store := db.NewStore(connPool)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server: " + err.Error())
	}
}
