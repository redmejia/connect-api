package router

import (
	"connect/api/handlers"
	"net/http"
)

func Routes(app *handlers.App) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/create-account", app.RegisterMyBusiness)
	mux.HandleFunc("/api/login", app.Signin)
	return mux
}
