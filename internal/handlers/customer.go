package handlers

import (
	"crm-app/internal/models"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Handler interface to handle HTTP requests
type Handler interface {
	GetCustomers(w http.ResponseWriter, r *http.Request)
	GetCustomerByID(w http.ResponseWriter, r *http.Request)
	CreateCustomer(w http.ResponseWriter, r *http.Request)
	UpdateCustomer(w http.ResponseWriter, r *http.Request)
	DeleteCustomer(w http.ResponseWriter, r *http.Request)
}

// CustomerHandler struct to implement the Handler interface
type CustomerHandler struct{}

var customers = []models.Customer{
	{ID: 1, Name: "Moh Test1 ", Role: "Manager", Email: "moh.test1@example.com", Phone: "123-456-7890", Contacted: true},
	{ID: 2, Name: "Jane Smith", Role: "Developer", Email: "jane.smith@example.com", Phone: "987-654-3210", Contacted: false},
	{ID: 3, Name: "Emily Johnson", Role: "Designer", Email: "emily.johnson@example.com", Phone: "555-555-5555", Contacted: true},
}

func (h *CustomerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (h *CustomerHandler) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
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

func (h *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var customer models.Customer
	if err := json.Unmarshal(body, &customer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	customer.ID = len(customers) + 1
	customers = append(customers, customer)
	json.NewEncoder(w).Encode(customer)
}

func (h *CustomerHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var updatedCustomer models.Customer
	if err := json.Unmarshal(body, &updatedCustomer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
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

func (h *CustomerHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
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
