package domain

type ListingRepositoryStub struct {
	listings []Listing
}

func (s ListingRepositoryStub) FindAll() ([]Listing, error) {
	return s.listings, nil
}

func NewListingRepositoryStub() ListingRepositoryStub {
	listings := []Listing{
		{"Pawamu", "22010", "3", "Kahawa", "Wendani", "false"},
		{"Ezulwini", "22010", "2", "Kahawa", "Sukari", "true"},
	}
	return ListingRepositoryStub{listings}
}
