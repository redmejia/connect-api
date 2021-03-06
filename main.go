package main

import (
	"connect/api/handlers"
	"connect/api/router"
	"connect/internal/database/postgresql"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	db, err := postgresql.Connection()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)

	app := handlers.App{
		Port:  8080,
		Info:  infoLog,
		Error: errorLog,
		DB: postgresql.DbPostgres{
			Db:    db,
			Info:  infoLog,
			Error: errorLog,
		},
		JwtKey: os.Getenv("JWT_KEY"),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.Port),
		Handler: router.Routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
