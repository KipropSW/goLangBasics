package domain

import (
	"server/dto"
	"server/errs"
)

type Account struct {
	AccountId   string  `json:"account_id"`
	ListingId   string  `json:"listing_id"`
	OpeningDate string  `json:"opening_date"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: a.AccountId}
}

type AccountRepository interface {
	Save(a Account) (*Account, *errs.AppError)
	FindBy(accountId string) (*Account, *errs.AppError)
	SaveTransaction(a Transaction) (*Transaction, *errs.AppError)
}

func (a Account) CanWithdraw(amount float64) bool {
	if a.Amount < amount {
		return false
	}
	return true
}
