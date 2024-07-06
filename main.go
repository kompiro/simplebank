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
	"github.com/techschool/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: " + err.Error())
	}
	pgxConfig, err := pgxpool.ParseConfig(config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: " + err.Error())
	}

	zlogger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()

	adapter := logadapter.NewLogger(zlogger)

	tracer := tracelog.TraceLog{
		Logger:   adapter,
		LogLevel: tracelog.LogLevelTrace,
	}
	pgxConfig.ConnConfig.Tracer = &tracer

	connPool, err := pgxpool.NewWithConfig(context.Background(), pgxConfig)
	if err != nil {
		log.Fatal("cannot connect to db: " + err.Error())
	}

	store := db.NewStore(connPool)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: " + err.Error())
	}
}
