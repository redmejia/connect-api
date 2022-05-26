package router

import (
	"connect/api/handlers"
	"net/http"
)

func Routes(app *handlers.App) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/create-account", app.RegisterMyBusiness)
	mux.HandleFunc("/api/login", app.Signin)
	mux.HandleFunc("/api/my-business", app.BusinessProfile)
	mux.HandleFunc("/api/my-business/deals", app.DealsByType)
	mux.HandleFunc("/api/my-business/deal", app.DealByIDs)
	return mux
}
