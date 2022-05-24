package handlers

import (
	"connect/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func (a *App) BusinessProfile(w http.ResponseWriter, r *http.Request) {

	// var fakeNewDeals models.Deal

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
		fakeBusiness := models.BusinessAccount{
			BusinessID:   53,
			BusinessName: "Connect the world is mine",
			BusinessType: "Software Development", // This is not type listed testing only two food and drink and agriculture more can be add
			Email:        "connect@mail.com",
			Founded:      1953,
			Password:     "****************",
		}
		var business models.BusinessAccount

		err := json.NewDecoder(r.Body).Decode(&business)
		if err != nil {
			a.Error.Println(err)
			return
		}

		a.Info.Println("before upadted ", fakeBusiness)

		if len(business.BusinessName) != 0 {
			fakeBusiness.BusinessName = business.BusinessName
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

// deals by type
func (a *App) DealsByType(w http.ResponseWriter, r *http.Request) {
	// http://localhost:8080/api/my-business/deals?type=fooddrink
	var dbDeals = make(map[string][]models.Deal)

	businessType := r.URL.Query().Get("type")

	if r.Method == http.MethodGet {
		a.Info.Println(businessType)

		deals := []models.Deal{
			models.Deal{
				DealID:          53,
				BusinessType:    "food & drink",
				ProductName:     "caps coffe",
				DealDescription: "I am sellig a box of coffee 16oz",
				DealStart:       time.Now(),
				// DealIsActive:    true,
				// Sold:            false,
				Price: 53.53,
			},
			models.Deal{
				DealID:          88,
				BusinessType:    "food & drink",
				ProductName:     "caps coffe",
				DealDescription: "I am sellig a box of coffee 16oz",
				DealStart:       time.Now(),
				// DealIsActive:    true,
				// Sold:            false,
				Price: 53.53,
			},
			models.Deal{
				DealID:          35,
				BusinessType:    "foo & drink",
				ProductName:     "test",
				DealDescription: "I am  bags 100pound and Semilla",
				DealStart:       time.Now(),
				// DealIsActive:    true,
				// Sold:            false,
				Price: 53.53,
			},
		}

		dbDeals["deals"] = deals

		data, err := json.Marshal(dbDeals)
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

func (a *App) DealByIdandType(w http.ResponseWriter, r *http.Request) {
	// http://localhost:8080/api/my-business/deal?deal-id=53&bus-type=fooddrink
	deals := []models.Deal{
		models.Deal{
			DealID:          53,
			BusinessType:    "food & drink",
			ProductName:     "caps coffe",
			DealDescription: "I am sellig a box of coffee 16oz",
			DealStart:       time.Now(),
			// DealIsActive:    true,
			// Sold:            false,
			Price: 53.53,
		},
		models.Deal{
			DealID:          88,
			BusinessType:    "food & drink",
			ProductName:     "caps coffe",
			DealDescription: "I am sellig a box of coffee 16oz",
			DealStart:       time.Now(),
			// DealIsActive:    true,
			// Sold:            false,
			Price: 53.53,
		},
		models.Deal{
			DealID:          35,
			BusinessID:      0,
			BusinessType:    "foo & drink",
			ProductName:     "test",
			DealDescription: "I am  bags 100pound and Semilla",
			DealStart:       time.Now(),
			IsActive:        models.ActiveDeals{Sold: false, DealIsActive: true},
			Price:           53.53,
		},
	}

	queryMap := r.URL.Query()

	id, _ := strconv.Atoi(queryMap["deal-id"][0])

	var deal = make(map[string]models.Deal)
	if r.Method == http.MethodGet {
		for _, v := range deals {
			if v.DealID == id {
				deal["deal"] = v
			}
		}
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		a.Error.Println("not found bad request")
		return
	}

	data, _ := json.Marshal(deal)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}
