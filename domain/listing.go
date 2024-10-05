package domain

import "server/errs"

type Listing struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Zipcode  string `json:"zipcode"`
	Location string `json:"location"`
	Ward     string `json:"ward"`
	Status   string `json:"status"`
}

type ListingRepository interface {
	FindAll(status string) ([]Listing, *errs.AppError)
	ById(string) (*Listing, *errs.AppError)
}
