package domain

import (
	"server/dto"
	"server/errs"
)

type Listing struct {
	Id       string `db:"listing_id"`
	Name     string
	Zipcode  string
	Location string
	Ward     string
	Status   string
}

func (l Listing) StatusAsText() string {
	statusAsText := "active"
	if l.Status == "false" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (l Listing) ToDto() dto.ListingResponse {
	return dto.ListingResponse{
		Name:     l.Name,
		Zipcode:  l.Zipcode,
		Location: l.Location,
		Ward:     l.Ward,
		Status:   l.StatusAsText(),
	}

}

type ListingRepository interface {
	FindAll(status string) ([]Listing, *errs.AppError)
	ById(string) (*Listing, *errs.AppError)
}
