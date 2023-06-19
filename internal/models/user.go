package models

type UserId uint
type ApiKey string

type User struct {
	ID     UserId `json:"id"`
	ApiKey ApiKey `json:"apiKey"`
}

type CreateUserDto struct {
	ApiKey ApiKey
}
