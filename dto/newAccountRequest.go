package dto

import (
	"server/errs"
	"strings"
)

type NewAccountRequest struct {
	ListingId   string  `json:"listing_id"`
	AccountType string  `json:"account_type"`
	Amount      float32 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("To open new account, you need to deposit at least 5000")
	}
	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return errs.NewValidationError("Account type should be  checking or saving")
	}
	return nil
}
