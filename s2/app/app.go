package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/greet", greet)
	router.HandleFunc("/greet", greet).Methods(http.MethodPost)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerId:[0-9]+}", getCustomer)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	err := http.ListenAndServe("localhost:8000", router)
	log.Fatal(err)
}
