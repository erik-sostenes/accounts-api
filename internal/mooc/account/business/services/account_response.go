package services

import "github.com/erik-sostenes/accounts-api/internal/mooc/account/business/domain"

// AccountResponse represent an DTO(Data Transfer Object)
type AccountResponse struct {
	AccountId       string
	AccountUserName string
	AccountName     string
	AccountLastName string
	AccountEmail    string
	AccountPassword string
	AccountCareer   string
	AccountIP       string
	AccountActive   bool
	AccountDetails  domain.Map
}
