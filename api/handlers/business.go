package handlers

import (
	"connect/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
)

// BusinessProfile For creating a profile, updating information and creating new deal.
func (a *App) BusinessProfile(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		// http://localhost:8080/api/my-business?bus-id=53

		businessId, err := strconv.Atoi(r.URL.Query().Get("bus-id"))
		if err != nil {
			a.Error.Println(err)
			return
		}

		business := a.DB.GetMyBusinessInfoById(businessId)

		data, _ := json.Marshal(business)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)

		a.Info.Println(string(data))

	case http.MethodPost:
		// http://localhost:8080/api/my-business

		var newDeal models.Deal

		json.NewDecoder(r.Body).Decode(&newDeal)

		ok := a.DB.CreateNewDeal(&newDeal)

		if ok {

			data, _ := json.Marshal(newDeal)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(data)

			a.Info.Println(string(data))
		}

	case http.MethodPatch:

		var businessUpdate models.BusinessAccount

		err := json.NewDecoder(r.Body).Decode(&businessUpdate)
		if err != nil {
			a.Error.Println(err)
			return
		}

		myBusinessInfo := a.DB.GetMyBusinessInfoById(businessUpdate.BusinessID)

		if len(businessUpdate.BusinessName) != 0 && businessUpdate.BusinessName != myBusinessInfo.BusinessName {
			myBusinessInfo.BusinessName = businessUpdate.BusinessName
		}

		if len(businessUpdate.Email) != 0 && businessUpdate.Email != myBusinessInfo.Email {
			myBusinessInfo.Email = businessUpdate.Email
		}

		if len(businessUpdate.BusinessType) != 0 && businessUpdate.BusinessType != myBusinessInfo.BusinessType {
			myBusinessInfo.BusinessType = businessUpdate.BusinessType
		}

		if businessUpdate.Founded > 0 && businessUpdate.Founded != myBusinessInfo.Founded {
			myBusinessInfo.Founded = businessUpdate.Founded
		}

		a.DB.UpdateProfile(myBusinessInfo)

		data, err := json.Marshal(myBusinessInfo)
		if err != nil {
			a.Error.Println(err)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(data)

		a.Info.Println("upadted ", string(data))

	default:
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		a.Info.Println("not found")
	}

}

// DealsByType
func (a *App) DealsByType(w http.ResponseWriter, r *http.Request) {
	// http://localhost:8080/api/my-business/deals?type=fooddrink

	businessType := r.URL.Query().Get("type")

	if r.Method == http.MethodGet {
		a.Info.Println(businessType)

		dealsTye := a.DB.GetDealsByType(businessType)

		data, err := json.Marshal(dealsTye)
		if err != nil {
			a.Error.Println(err)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
		a.Info.Println(string(data))

	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		a.Error.Println("not found ")
	}
}

func (a *App) DealByIDs(w http.ResponseWriter, r *http.Request) {
	// http://localhost:8080/api/my-business/deal?deal-id=17&bus-id=3

	queryMap := r.URL.Query()

	did, _ := strconv.Atoi(queryMap["deal-id"][0])
	bid, _ := strconv.Atoi(queryMap["bus-id"][0])

	if r.Method == http.MethodGet {
		deal := a.DB.GetDealsByIDs(did, bid)

		data, _ := json.Marshal(deal)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		a.Error.Println("not found bad request")
		return
	}

}

// DeleteByIDs
func (a *App) DeleteDeal(w http.ResponseWriter, r *http.Request) {
	// http://localhost/api/my-business/del/deal?deal-id=15&bus-id=1
	queryMap := r.URL.Query()
	did, _ := strconv.Atoi(queryMap["deal-id"][0])
	bid, _ := strconv.Atoi(queryMap["bus-id"][0])

	if r.Method == http.MethodDelete {
		ok := a.DB.DeleteDealByIDs(did, bid)

		if ok {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			a.Info.Println("deal was deleted")
		} else {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			a.Error.Println("not found bad request")
			return
		}
	}
}
