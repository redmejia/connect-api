package models

// BusinessAccount creating account more require infomation could be add here.
type BusinessAccount struct {
	BussinesName string `json:"bussiness_name"`
	Email        string `json:"email"` // for login
	Founded      int    `json:"founded"`
	Password     string `json:"-"`
}

// LogIn only email and password.
type LogIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
