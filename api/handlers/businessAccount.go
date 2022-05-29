package handlers

import (
	"connect/internal/models"
	"encoding/json"
	"net/http"
)

// RegisterMyBusiness
func (a *App) RegisterMyBusiness(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		var myBusinessAccount models.BusinessAccount

		err := json.NewDecoder(r.Body).Decode(&myBusinessAccount)
		if err != nil {
			a.Error.Fatal(err)
			return
		}

		ok := a.DB.RegisterMyBusiness(&myBusinessAccount)
		if ok {

			data, err := json.Marshal(&myBusinessAccount)
			if err != nil {
				a.Error.Println(err)
				return
			}

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			// w.Write(data)

			a.Info.Println(string(data))
		}
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

	}

}
