package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ravikrs/go-rest-banking/s3/dto"
	"github.com/ravikrs/go-rest-banking/s3/errs"
	"github.com/ravikrs/go-rest-banking/s3/service"
	"net/http"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	statusQueryParam := r.URL.Query().Get("status")
	var customers []dto.CustomerResponse
	var err *errs.AppError
	if statusQueryParam != "" {
		customers, err = ch.service.GetAllCustomerByStatus(statusQueryParam)
	} else {
		customers, err = ch.service.GetAllCustomer()
	}
	if err != nil {
		writeResponse(w, err.Code, err)
		return
	}
	writeResponse(w, http.StatusOK, customers)
}

func (ch CustomerHandler) getCustomerByID(w http.ResponseWriter, r *http.Request) {
	customerId := mux.Vars(r)["customer_id"]
	customer, err := ch.service.GetCustomerByID(customerId)
	if err != nil {
		writeResponse(w, err.Code, customer)
		return
	}
	writeResponse(w, http.StatusOK, customer)
}

func writeResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "UP")
}
