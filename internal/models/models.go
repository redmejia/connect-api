package models

import "time"

// BusinessAccount creating account more require infomation could be add here.
type BusinessAccount struct {
	BusinessID   int    `json:"business_id"`
	BusinessName string `json:"business_name"`
	BusinessType string `json:"business_type"`
	Email        string `json:"email"`
	Founded      int    `json:"founded"`
	Password     string `json:"-"`
}

// LogIn only email and password.
type LogIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ActiveDeal
type ActiveDeals struct {
	DealID       int  `json:"deal_id"`
	BusinessID   int  `json:"business_id"`
	DealIsActive bool `json:"deal_is_active"`
	Sold         bool `json:"sold"`
}

// Deal
type Deal struct {
	DealID          int         `json:"deal_id"`
	BusinessID      int         `json:"business_id"`
	BusinessName    string      `json:"business_name"`
	BusinessType    string      `json:"business_type"`
	ProductName     string      `json:"product_name"`
	DealDescription string      `json:"deal_desciption"`
	DealStart       time.Time   `json:"deal_start"` // change to created at
	IsActive        ActiveDeals `json:"is_active"`
	Price           float64     `json:"price"`
}
