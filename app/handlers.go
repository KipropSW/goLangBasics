package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"server/service"
)

type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zipcode" xml:"zipcode"`
}

//func greet(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "Hello World")
//}

type ListingHandlers struct {
	service service.ListingService
}

func (ch *ListingHandlers) getAllListing(w http.ResponseWriter, r *http.Request) {
	//customers := []Customer{
	//	{"Kiprop", "Nairobi", "200001"},
	//	{"Silot", "Kericho", "300001"},
	//	{"Bett", "Oleng", "200001"},
	//}
	status := r.URL.Query().Get("status")
	listings, err := ch.service.GetAllListing(status)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, err.AsMessage())
	}
	writeResponse(w, http.StatusOK, listings)
	//
	//if r.Header.Get("Content-Type") == "application/xml" {
	//	w.Header().Add("Content-Type", "application/xml")
	//	xml.NewEncoder(w).Encode(listings)
	//} else {
	//	w.Header().Add("Content-Type", "application/json")
	//	json.NewEncoder(w).Encode(listings)
	//}
}

func (ch *ListingHandlers) getListing(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["listing_id"]
	listing, err := ch.service.GetListing(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, listing)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

//func getAllCustomers(w http.ResponseWriter, r *http.Request) {
//	customers := []Customer{
//		{"Kiprop", "Nairobi", "200001"},
//		{"Silot", "Kericho", "300001"},
//		{"Bett", "Oleng", "200001"},
//	}
//
//	if r.Header.Get("Content-Type") == "application/xml" {
//		w.Header().Add("Content-Type", "application/xml")
//		xml.NewEncoder(w).Encode(customers)
//	} else {
//		w.Header().Add("Content-Type", "application/json")
//		json.NewEncoder(w).Encode(customers)
//	}
//}
