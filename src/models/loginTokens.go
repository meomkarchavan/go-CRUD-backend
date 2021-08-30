package models

var LoginTokens = map[string]*LoginToken{}

type LoginToken struct {
	UserId        string
	Refresh_token string
	Access_token  string
}
