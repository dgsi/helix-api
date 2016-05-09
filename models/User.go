package models

import (
	"time"
)

type User struct {
	Id int `json:"id"`
	Clientid  string `json:"client_id"`
	Username  string `json:"username"`	
	Password string `json:"password"`
	Companyid  string `json:"company_id"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
	Status string `json:"status"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}