package models

type AuthenticatedUser struct {
	Id int `json:"id"`
	Clientid  string `json:"client_id"`
	Username  string `json:"username"`	
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Companyid  string `json:"company_id"`
	Token string `json:"token"`
}