package handlers

import (
	"connect/internal/models"
	"connect/internal/utils"
	"encoding/json"
	"net/http"
)

func (a *App) Signin(w http.ResponseWriter, r *http.Request) {

	var business models.LogIn

	err := json.NewDecoder(r.Body).Decode(&business)
	if err != nil {
		a.Error.Println(err)
	}

	token, err := utils.GenToken(business.Email)
	if err != nil {
		a.Error.Println(err)
		return
	}

	var success = struct {
		IsLogedin bool   `json:"is_login"`
		Token     string `json:"token"`
	}{
		IsLogedin: true,
		Token:     token,
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "bus_jwt",
		Value:    token,
		HttpOnly: true,
	})

	data, err := json.Marshal(success)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}
