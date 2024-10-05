package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"server/domain"
	"server/service"
)

func Start() {

	// multiplexer
	//mux := http.NewServeMux()
	router := mux.NewRouter()

	//wiring
	ls := ListingHandlers{service: service.NewDefaultListingService(domain.NewListingRepositoryDb())}

	// define routes
	//router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/listing", ls.getAllListing).Methods(http.MethodGet)
	router.HandleFunc("/listing/{listing_id:[0-9]+}", ls.getListing).Methods(http.MethodGet)

	//router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getAllCustomer).Methods(http.MethodGet)

	//starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")
}

func getAllCustomer(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprint(writer, vars["customer_id"])
}
