package handlers

import (
	"connect/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func (a *App) BusinessProfile(w http.ResponseWriter, r *http.Request) {
	// test data model
	fakeBusiness := models.BusinessAccount{
		BusinessID:   53,
		BussinesName: "Connect the world is mine",
		Email:        "connect@mail.com",
		Founded:      1953,
		Password:     "****************",
	}

	var fakeNewDeals models.Deal

	switch r.Method {
	case http.MethodGet:
		// for profile information
		// http://localhost:8080/api/my-business?bus-id=53
		id, err := strconv.Atoi(r.URL.Query().Get("bus-id"))
		if err != nil {
			a.Error.Println(err)
			return
		}

		a.Info.Println("your business-id ", id)

		if id == fakeBusiness.BusinessID {

			data, _ := json.Marshal(fakeBusiness)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(data)

			a.Info.Println(string(data))

		}

	case http.MethodPost:
		var newDeal models.Deal

		json.NewDecoder(r.Body).Decode(&newDeal)

		fakeNewDeals.ProductName = newDeal.ProductName
		fakeNewDeals.DealDescription = newDeal.DealDescription
		fakeNewDeals.DealStart = time.Now()
		fakeNewDeals.DealIsActive = false
		fakeNewDeals.Sold = true
		fakeNewDeals.Price = newDeal.Price

		data, _ := json.Marshal(fakeNewDeals)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)

		a.Info.Println(string(data))

	case http.MethodPatch:
		var business models.BusinessAccount

		err := json.NewDecoder(r.Body).Decode(&business)
		if err != nil {
			a.Error.Println(err)
			return
		}

		a.Info.Println("before upadted ", fakeBusiness)

		if len(business.BussinesName) != 0 {
			fakeBusiness.BussinesName = business.BussinesName
		}

		if len(business.Email) != 0 {
			fakeBusiness.Email = business.Email
		}

		if business.Founded > 0 && business.Founded != fakeBusiness.Founded {
			fakeBusiness.Founded = business.Founded
		}

		data, err := json.Marshal(fakeBusiness)
		if err != nil {
			a.Error.Println(err)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)

		a.Info.Println("upadted ", string(data))

	default:
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		a.Info.Println("not found")
	}

}
