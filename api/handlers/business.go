package handlers

import (
	"connect/internal/models"
	"connect/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// BusinessProfile For creating a profile, updating information and creating new deal.
func (a *App) BusinessProfile(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		// http://localhost:8080/api/my/business?bus-id=53

		businessId, err := strconv.Atoi(r.URL.Query().Get("bus-id"))
		if err != nil {
			a.Error.Println(err)
			return
		}

		business := a.DB.GetMyBusinessInfoById(businessId)

		err = utils.WriteJson(w, http.StatusOK, "myBusiness", business)
		if err != nil {
			a.Error.Println(err)
			return
		}

		a.Info.Println(*business)

	case http.MethodPost:
		// http://localhost:8080/api/my/business

		var newDeal models.Deal

		json.NewDecoder(r.Body).Decode(&newDeal)

		dealId, ok := a.DB.CreateNewDeal(&newDeal)

		if ok {

			newDeal.DealID = dealId
			newDeal.DealStart = time.Now()

			err := utils.WriteJson(w, http.StatusCreated, "myDeal", newDeal)
			if err != nil {
				a.Error.Println(err)
				return
			}

			a.Info.Println(newDeal)
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

		err = utils.WriteJson(w, http.StatusCreated, "myBusiness", myBusinessInfo)
		if err != nil {
			a.Error.Println(err)
			return
		}

		a.Info.Println("upadted ", myBusinessInfo)

	case http.MethodOptions:
		return

	default:
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		a.Info.Println("not found")
	}

}

// DealsByType
func (a *App) DealsByType(w http.ResponseWriter, r *http.Request) {
	// http://localhost:8080/api/my/business/deals?type=fooddrink

	if r.Method == http.MethodOptions {
		log.Println("you are here now")
		return
	} else {

		businessType := r.URL.Query().Get("type")

		dealsType := a.DB.GetDealsByType(businessType)

		err := utils.WriteJson(w, http.StatusOK, "deals", dealsType)
		if err != nil {
			a.Error.Println(err)
			return
		}

		a.Info.Println(*dealsType)
	}
	// switch r.Method {
	// case http.MethodGet:

	// 	businessType := r.URL.Query().Get("type")

	// 	dealsType := a.DB.GetDealsByType(businessType)

	// 	w.Header().Add("Content-Type", "application/json")
	// 	err := json.NewEncoder(w).Encode(dealsType)
	// 	// err := utils.WriteJson(w, http.StatusOK, "deals", dealsType)
	// 	if err != nil {
	// 		a.Error.Println(err)
	// 		return
	// 	}

	// 	a.Info.Println(*dealsType)
	// case http.MethodOptions:
	// 	log.Println("OPTIONBS")
	// 	return
	// default:
	// 	return
	// }

	// else {
	// 	w.Header().Add("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	a.Error.Println("not found ")
	// }
}

func (a *App) DealByIDs(w http.ResponseWriter, r *http.Request) {
	// http://localhost:8080/api/my/business/deal?deal-id=17&bus-id=3

	if r.Method == http.MethodGet {
		queryMap := r.URL.Query()

		did, _ := strconv.Atoi(queryMap["deal-id"][0])
		bid, _ := strconv.Atoi(queryMap["bus-id"][0])

		deal := a.DB.GetDealsByIDs(did, bid)

		err := utils.WriteJson(w, http.StatusOK, "deal", deal)
		if err != nil {
			a.Error.Println(err)
			return
		}

		a.Info.Println(deal)

	} else if r.Method == http.MethodOptions {
		return

	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		a.Error.Println("not found bad request")
		return
	}

}

// Message
type Message struct {
	Error bool   `json:"error"`
	Msg   string `json:"msg"`
}

// DeleteByIDs
func (a *App) DeleteDeal(w http.ResponseWriter, r *http.Request) {
	// http://localhost/api/my/business/del/deal?deal-id=15&bus-id=1

	if r.Method == http.MethodDelete {
		queryMap := r.URL.Query()
		did, _ := strconv.Atoi(queryMap["deal-id"][0])
		bid, _ := strconv.Atoi(queryMap["bus-id"][0])

		ok := a.DB.DeleteDealByIDs(did, bid)

		if ok {
			var message Message

			message.Error = false
			message.Msg = fmt.Sprintf("Deal No %d was deleted ", did)

			err := utils.WriteJson(w, http.StatusAccepted, "deal", message)
			if err != nil {
				a.Error.Println(err)
				return
			}

			a.Info.Println(message)
		} else if r.Method == http.MethodOptions {
			return
		} else {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			a.Error.Println("not found bad request")
			return
		}
	}
}

// DealUpdate update deal status if owner wants to sell same product if was not deleted
func (a *App) DealUpdate(w http.ResponseWriter, r *http.Request) {
	// http://localhost:8080/api/my/business/deal/stat

	if r.Method == http.MethodPatch {
		var dealUpdate models.ActiveDeals

		json.NewDecoder(r.Body).Decode(&dealUpdate)
		ok := a.DB.UpdateDealStatus(&dealUpdate)

		if ok {

			var message Message
			message.Error = false
			message.Msg = fmt.Sprintf("Update deal status where ids No %d - %d", dealUpdate.DealID, dealUpdate.BusinessID)

			err := utils.WriteJson(w, http.StatusAccepted, "deal", message)
			if err != nil {
				a.Error.Println(err)
				return
			}

			a.Info.Println(dealUpdate)

		} else {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			a.Error.Println("not found bad request")
			return
		}
	}

}
