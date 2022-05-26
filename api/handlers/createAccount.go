package handlers

import (
	"connect/internal/models"
	"encoding/json"
	"net/http"
)

// test map fake database
var DBs = make(map[string]models.BusinessAccount)

func (a *App) RegisterMyBusiness(w http.ResponseWriter, r *http.Request) {
	var myBusinessAccount models.BusinessAccount
	a.Info.Println(myBusinessAccount)

	err := json.NewDecoder(r.Body).Decode(&myBusinessAccount)
	if err != nil {
		a.Error.Fatal(err)
		return
	}

	DBs["my-business"] = myBusinessAccount

	data, err := json.Marshal(DBs)
	if err != nil {
		a.Error.Println(err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)

	a.Info.Println(string(data))
}
