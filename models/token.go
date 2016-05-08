package models

type AuthenticatedUser struct {
	Token string `json:"token"`
}