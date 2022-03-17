package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ravikrs/go-rest-banking/s3/domain"
	"github.com/ravikrs/go-rest-banking/s3/service"
	"log"
	"net/http"
	"os"
)

func Start() {
	r := mux.NewRouter()
	//wiring
	//handler := CustomerHandler{service.NewDefaultCustomerService(domain.NewCustomerRepositoryStub())}
	handler := CustomerHandler{service.NewDefaultCustomerService(domain.NewCustomerRepositoryDB())}

	//define routes
	r.HandleFunc("/customers", handler.getAllCustomers).Methods(http.MethodGet)
	r.HandleFunc("/customers/{customer_id:[0-9]+}", handler.getCustomerByID).Methods(http.MethodGet)
	r.HandleFunc("/health", health)
	port := os.Getenv("SERVER_PORT")
	address := os.Getenv("SERVER_ADDRESS")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), r))
}
