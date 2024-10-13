package dto

type ListingResponse struct {
	//Id       string `json:"id"`
	Name     string `json:"name"`
	Zipcode  string `json:"zipcode"`
	Location string `json:"location"`
	Ward     string `json:"ward"`
	Status   string `json:"status"`
}
