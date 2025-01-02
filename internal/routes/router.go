package routes

import (
	"crm-app/internal/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/customers", handlers.GetCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", handlers.GetCustomerByID).Methods("GET")
	router.HandleFunc("/customers", handlers.CreateCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", handlers.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", handlers.DeleteCustomer).Methods("DELETE")
	return router
}
