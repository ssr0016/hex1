package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ssr0016/goHex/service"
)

type customerHandler struct {
	custSvc service.CustomerService
}

func NewCustomerHandler(custSvc service.CustomerService) customerHandler {
	return customerHandler{
		custSvc: custSvc,
	}
}

func (h customerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.custSvc.GetCustomers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (h customerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])

	customer, err := h.custSvc.GetCustomer(customerID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
