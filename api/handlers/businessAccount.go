package handlers

import (
	"connect/internal/models"
	"connect/utils"
	"encoding/json"
	"net/http"
)

// RegisterMyBusiness
func (a *App) RegisterMyBusiness(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		var myBusinessAccount models.BusinessAccount

		err := json.NewDecoder(r.Body).Decode(&myBusinessAccount)
		if err != nil {
			a.Error.Println(err)
		}

		a.Info.Println(myBusinessAccount)

		basicInfo, ok := a.DB.RegisterMyBusiness(&myBusinessAccount)
		if ok {

			token, err := utils.GenToken(myBusinessAccount.Email)
			if err != nil {
				a.Error.Println(err)
				return
			}

			var success = struct {
				BusinessId   int    `json:"business_id"`
				BusinessName string `json:"business_name"`
				BusinessType string `json:"business_type"`
				IsAuth       bool   `json:"is_auth"`
				Token        string `json:"token"`
			}{
				BusinessId:   basicInfo.BusinessID,
				BusinessName: basicInfo.BusinessName,
				BusinessType: basicInfo.BusinessType,
				IsAuth:       true,
				Token:        token,
			}
			// var success = struct {
			// 	BusinessId int    `json:"business_id"`
			// 	IsAuth     bool   `json:"is_auth"`
			// 	Token      string `json:"token"`
			// }{
			// 	BusinessId: businessID,
			// 	IsAuth:     true,
			// 	Token:      token,
			// }

			err = utils.WriteJson(w, http.StatusOK, "success", success)
			if err != nil {
				a.Error.Println(err)
				return
			}
			a.Info.Println(success)
		}
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

	}

}
