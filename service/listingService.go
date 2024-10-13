package service

import (
	"server/domain"
	"server/dto"
	"server/errs"
)

type ListingService interface {
	GetAllListing(string) ([]domain.Listing, *errs.AppError)
	GetListing(string) (*dto.ListingResponse, *errs.AppError)
}

type DefaultListingService struct {
	repo domain.ListingRepository
}

func (s DefaultListingService) GetAllListing(status string) ([]domain.Listing, *errs.AppError) {
	if status == "active" {
		status = "true"
	} else if status == "inactive" {
		status = "false"
	} else {
		status = ""
	}
	return s.repo.FindAll(status)
}

func (s DefaultListingService) GetListing(id string) (*dto.ListingResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()

	//response := dto.ListingResponse{
	//	Id:       c.Id,
	//	Name:     c.Name,
	//	Zipcode:  c.Zipcode,
	//	Location: c.Location,
	//	Ward:     c.Ward,
	//	Status:   c.Status,
	//}
	return &response, nil
}

func NewDefaultListingService(repo domain.ListingRepository) DefaultListingService {
	return DefaultListingService{repo: repo}
}
