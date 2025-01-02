package handlers

import (
	"crm-app/internal/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var customers = []models.Customer{}

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for _, customer := range customers {
		if customer.ID == id {
			json.NewEncoder(w).Encode(customer)
			return
		}
	}
	http.Error(w, "Customer not found", http.StatusNotFound)
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customer models.Customer
	json.NewDecoder(r.Body).Decode(&customer)
	customer.ID = len(customers) + 1
	customers = append(customers, customer)
	json.NewEncoder(w).Encode(customer)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var updatedCustomer models.Customer
	json.NewDecoder(r.Body).Decode(&updatedCustomer)
	for i, customer := range customers {
		if customer.ID == id {
			customers[i] = updatedCustomer
			customers[i].ID = id
			json.NewEncoder(w).Encode(customers[i])
			return
		}
	}
	http.Error(w, "Customer not found", http.StatusNotFound)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for i, customer := range customers {
		if customer.ID == id {
			customers = append(customers[:i], customers[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Customer not found", http.StatusNotFound)
}
