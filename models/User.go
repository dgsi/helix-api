package models

import (
	"time"
)

type User struct {
	Id int `json:"id"`
	ClientId  string `json:"client_id"`
	Username  string `json:"username"`	
	Password string `json:"password"`
	CompanyId  string `json:"company_id"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}