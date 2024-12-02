package handler

import (
	"github.com/nrmnqdds/vtech-qms-be/config/logger"
	"github.com/nrmnqdds/vtech-qms-be/db/store"

	"github.com/jackc/pgx/v5"
)

type App struct {
	ConnPool *pgx.Conn
	Queries  *store.Queries
	Logger   *logger.AppLogger
}

func NewApp(connPool *pgx.Conn) *App {
	return &App{
		ConnPool: connPool,
		Queries:  store.New(connPool),
		Logger:   logger.New(),
	}
}
