package handlers

import (
	"connect/internal/database/postgresql"
	"log"
)

type App struct {
	Port  int
	Info  *log.Logger
	Error *log.Logger
	DB    postgresql.DbPostgres
}
