package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

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
		return
	}

	err = util.SetupLogger(pgxConfig)
	if err != nil {
		log.Fatal("cannot setup logger: " + err.Error())
	}

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
