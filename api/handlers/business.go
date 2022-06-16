package handlers

import (
	"connect/internal/models"
	"connect/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// BusinessProfile For creating a profile, updating information and creating new deal.
func (a *App) BusinessProfile(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		// Get the business information
		// http://localhost:8080/api/my/business?bus-id=53

		businessId, err := strconv.Atoi(r.URL.Query().Get("bus-id"))
		if err != nil {
			a.Error.Println(err)
			return
		}

		myBusiness := a.DB.GetMyBusinessInfo(businessId)

		err = utils.WriteJson(w, http.StatusOK, "myBusiness", *myBusiness)
		if err != nil {
			a.Error.Println(err)
			return
		}

		a.Info.Println("YOU ARE HERE", *myBusiness)

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

		myBusiness := a.DB.GetMyBusinessInfo(myBusinessInfo.BusinessID)

		err = utils.WriteJson(w, http.StatusCreated, "myBusiness", *myBusiness)
		if err != nil {
			a.Error.Println(err)
			return
		}

		a.Info.Println("upadted ", myBusinessInfo)
	default:
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		a.Info.Println("not found")
	}

}

// MyDealsOrOffer
func (a *App) MyDealsOrOffer(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		// http://localhost:8080/api/my/business/my/deals?bus-id=1
		businessId, _ := strconv.Atoi(r.URL.Query().Get("bus-id"))

		myDealsOrOffer := a.DB.GetMyDealOrOffer(businessId)

		if len(*myDealsOrOffer) == 0 {

			emptydeal := []models.Deal{}
			err := utils.WriteJson(w, http.StatusOK, "myDeals", emptydeal)
			if err != nil {
				a.Error.Println(err)
			}

			a.Info.Println(*myDealsOrOffer)

		} else {

			err := utils.WriteJson(w, http.StatusOK, "myDeals", myDealsOrOffer)
			if err != nil {
				a.Error.Println(err)
			}

			a.Info.Println(*myDealsOrOffer)
		}

	case http.MethodPatch:
		var myDeal models.Deal

		json.NewDecoder(r.Body).Decode(&myDeal)

		deal := a.DB.GetDealsByIDs(myDeal.DealID, myDeal.BusinessID)

		if len(myDeal.ProductName) != 0 && myDeal.ProductName != deal.ProductName {
			deal.ProductName = myDeal.ProductName
		}

		if len(myDeal.DealDescription) != 0 && myDeal.DealDescription != deal.DealDescription {
			deal.DealDescription = myDeal.DealDescription
		}

		if myDeal.Price > 0 && myDeal.Price != deal.Price {
			deal.Price = myDeal.Price

		}

		ok := a.DB.UpdateMyDealOrOffer(&deal)
		if ok {
			err := utils.WriteJson(w, http.StatusAccepted, "myDeal", deal)
			if err != nil {
				a.Error.Println(err)
			}
			a.Info.Println(deal)
		}

	default:
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		a.Info.Println("not found")
	}
}

// DealsByType
func (a *App) DealsByType(w http.ResponseWriter, r *http.Request) {
	// http://localhost:8080/api/my/business/deals?type=fooddrink

	switch r.Method {
	case http.MethodGet:

		businessType := r.URL.Query().Get("type")

		dealsType := a.DB.GetDealsByType(businessType)

		err := utils.WriteJson(w, http.StatusOK, "deals", dealsType)
		if err != nil {
			a.Error.Println(err)
			return
		}

		a.Info.Println(*dealsType)
	default:
		return
	}

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

			// err := utils.WriteJson(w, http.StatusAccepted, "deal", message)
			// if err != nil {
			// 	a.Error.Println(err)
			// 	return
			// }

			a.Info.Println(message)
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
