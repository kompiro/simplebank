package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/techschool/simplebank/util"
)

var testStore Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot connect to config: " + err.Error())
	}

	pgxConfig, err := pgxpool.ParseConfig(config.DBSource)
	if err != nil {
		log.Fatal("cannot setup logger: " + err.Error())
	}

	connPool, err := pgxpool.NewWithConfig(context.Background(), pgxConfig)
	if err != nil {
		log.Fatal("cannot connect to db: " + err.Error())
	}

	testStore = NewStore(connPool)
	os.Exit(m.Run())
}
