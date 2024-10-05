package service

import (
	"server/domain"
	"server/errs"
)

type ListingService interface {
	GetAllListing(string) ([]domain.Listing, *errs.AppError)
	GetListing(string) (*domain.Listing, *errs.AppError)
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

func (s DefaultListingService) GetListing(id string) (*domain.Listing, *errs.AppError) {
	return s.repo.ById(id)
}

func NewDefaultListingService(repo domain.ListingRepository) DefaultListingService {
	return DefaultListingService{repo: repo}
}
