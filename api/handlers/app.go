package handlers

import "log"

type App struct {
	Port  int
	Info  *log.Logger
	Error *log.Logger
}
