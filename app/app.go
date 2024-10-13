package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
	"server/domain"
	"server/service"
	"time"
)

func Start() {

	//SanityCheck()

	// multiplexer
	//mux := http.NewServeMux()
	router := mux.NewRouter()

	dbClient := getDbClient()
	listingRepositoryDb := domain.NewListingRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	//wiring
	ls := ListingHandlers{service: service.NewDefaultListingService(listingRepositoryDb)}
	ah := AccountHandler{service: service.NewAccountService(accountRepositoryDb)}

	// define routes
	//router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/listing", ls.getAllListing).Methods(http.MethodGet)
	router.HandleFunc("/listing/{listing_id:[0-9]+}", ls.getListing).Methods(http.MethodGet)
	router.HandleFunc("/listing/{listing_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/listing/{listing_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).
		Methods(http.MethodPost)

	//router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/", getAllCustomer).Methods(http.MethodGet)

	//starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func getDbClient() *sqlx.DB {
	//dbUser := os.Getenv("DB_USER")
	//dbPassword := os.Getenv("DB_PASSWORD")
	//dbAddress := os.Getenv("DB_ADDRESS")
	//dbPort := os.Getenv("DB_PORT")
	//dbName := os.Getenv("DB_NAME")
	//dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbAddress, dbPort, dbName)

	//client, err := sqlx.Open("mysql", dataSource)

	client, err := sqlx.Open("mysql", "root:Silot777@@tcp(localhost:3306)/listing?charset=utf8")
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}

func SanityCheck() {
	if os.Getenv("DB_ADDRESS") == "" || os.Getenv("DB_PORT") == "" {
		log.Fatal("Environment variables SERVER_ADDRESS and SERVER_PORT must be set.")
	}
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")
}

func getAllCustomer(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprint(writer, vars["customer_id"])
}
