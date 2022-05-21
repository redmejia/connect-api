package handlers

import (
	"connect/internal/models"
	"encoding/json"
	"net/http"
)

func (a *App) RegisterMyBusiness(w http.ResponseWriter, r *http.Request) {
	var myBusinessAccount models.BusinessAccount

	err := json.NewDecoder(r.Body).Decode(&myBusinessAccount)
	if err != nil {
		a.Error.Fatal(err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	a.Info.Println("Recived ", myBusinessAccount)
}
