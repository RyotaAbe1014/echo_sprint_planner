package models

type Token struct {
	AccessToken  string  `json:"access"`
	RefreshToken *string `json:"refresh"`
}
