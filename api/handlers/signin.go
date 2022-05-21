package handlers

import (
	"connect/internal/models"
	"encoding/json"
	"net/http"
)

// for testing
var registeredDB models.LogIn

func (a *App) Signin(w http.ResponseWriter, r *http.Request) {
	registeredDB.Email = "connect@mail.com"
	registeredDB.Password = "theking"

	var businessRegistered models.LogIn

	err := json.NewDecoder(r.Body).Decode(&businessRegistered)
	if err != nil {
		a.Error.Println(err)
	}

	ok := false

	if businessRegistered.Email == registeredDB.Email && businessRegistered.Password == registeredDB.Password {
		ok = true
	}

	if ok {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Succes"))
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("No match"))
	}

}
