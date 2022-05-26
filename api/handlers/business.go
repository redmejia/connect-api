package handlers

import (
	"connect/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

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

// deals by type
func (a *App) DealsByType(w http.ResponseWriter, r *http.Request) {
	// http://localhost:8080/api/my-business/deals?type=fooddrink
	// var dbDeals = make(map[string][]models.Deal)

	businessType := r.URL.Query().Get("type")

	if r.Method == http.MethodGet {
		a.Info.Println(businessType)

		// deals := []models.Deal{
		// 	models.Deal{
		// 		DealID:          53,
		// 		BusinessType:    "food & drink",
		// 		ProductName:     "caps coffe",
		// 		DealDescription: "I am sellig a box of coffee 16oz",
		// 		DealStart:       time.Now(),
		// 		// DealIsActive:    true,
		// 		// Sold:            false,
		// 		Price: 53.53,
		// 	},
		// 	models.Deal{
		// 		DealID:          88,
		// 		BusinessType:    "food & drink",
		// 		ProductName:     "caps coffe",
		// 		DealDescription: "I am sellig a box of coffee 16oz",
		// 		DealStart:       time.Now(),
		// 		// DealIsActive:    true,
		// 		// Sold:            false,
		// 		Price: 53.53,
		// 	},
		// 	models.Deal{
		// 		DealID:          35,
		// 		BusinessType:    "foo & drink",
		// 		ProductName:     "test",
		// 		DealDescription: "I am  bags 100pound and Semilla",
		// 		DealStart:       time.Now(),
		// 		// DealIsActive:    true,
		// 		// Sold:            false,
		// 		Price: 53.53,
		// 	},
		// }

		// dbDeals["deals"] = deals
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
