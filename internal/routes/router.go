package routes

import (
	"crm-app/internal/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()
	handler := &handlers.CustomerHandler{}
	router.HandleFunc("/customers", handler.GetCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", handler.GetCustomerByID).Methods("GET")
	router.HandleFunc("/customers", handler.CreateCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", handler.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", handler.DeleteCustomer).Methods("DELETE")
	return router
}
