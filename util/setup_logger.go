package util

import (
	"os"

	logadapter "github.com/jackc/pgx-zerolog"
	"github.com/rs/zerolog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
)

func SetupLogger(pgxConfig *pgxpool.Config) (err error) {
	zlogger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()

	adapter := logadapter.NewLogger(zlogger)

	tracer := tracelog.TraceLog{
		Logger:   adapter,
		LogLevel: tracelog.LogLevelTrace,
	}
	pgxConfig.ConnConfig.Tracer = &tracer
	return
}
