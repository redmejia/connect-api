package router

import (
	"connect/api/handlers"
	"connect/api/middleware"
	"net/http"
)

func Routes(app *handlers.App) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/create/account", app.RegisterMyBusiness)
	mux.HandleFunc("/api/login", app.Signin)

	mux.HandleFunc("/api/my/business", middleware.IsAuthorizationToken(app.BusinessProfile))
	mux.HandleFunc("/api/my/business/my/deals", middleware.IsAuthorizationToken(app.MyDealsOrOffer))

	mux.HandleFunc("/api/my/business/deals", middleware.IsAuthorizationToken(app.DealsByType))
	mux.HandleFunc("/api/my/business/deal", middleware.IsAuthorizationToken(app.DealByIDs))

	mux.Handle("/api/my/business/del/deal", middleware.IsAuthorizationToken(app.DeleteDeal))
	mux.Handle("/api/my/business/deal/stat", middleware.IsAuthorizationToken(app.DealUpdate))

	return middleware.Cors(mux)
}
