package handlers

import (
	"connect/internal/models"
	"connect/internal/utils"
	"encoding/json"
	"net/http"
)

func (a *App) Signin(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		var business models.LogIn

		err := json.NewDecoder(r.Body).Decode(&business)
		if err != nil {
			a.Error.Println(err)
		}

		businessAuthInfo := a.DB.GetAuthInfo(business.Email)

		ok, err := utils.ComparePassword(businessAuthInfo.Password, business.Password)
		if err != nil {
			a.Error.Println(err)
			return
		}

		if ok {
			token, err := utils.GenToken(business.Email)
			if err != nil {
				a.Error.Println(err)
				return
			}

			var success = struct {
				BusinessId int    `json:"business_id"`
				IsAuth     bool   `json:"is_auth"`
				Token      string `json:"token"`
			}{
				BusinessId: businessAuthInfo.BusinessID,
				IsAuth:     true,
				Token:      token,
			}

			// http.SetCookie(w, &http.Cookie{
			// 	Name:     "bus_jwt",
			// 	Value:    token,
			// 	HttpOnly: true,
			// })

			data, err := json.Marshal(success)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		}
	}

}
