package router

import (
	"connect/api/handlers"
	"net/http"
)

func Routes(app *handlers.App) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.SayHello)
	mux.HandleFunc("/api/create-account", app.RegisterMyBusiness)
	return mux
}
