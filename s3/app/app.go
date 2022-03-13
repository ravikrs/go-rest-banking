package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/greet", greet)

	err := http.ListenAndServe("localhost:8000", router)
	log.Fatal(err)
}