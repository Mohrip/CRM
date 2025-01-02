package routes

import (
	"crm-app/internal/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()
	handler := &handlers.CustomerHandler{}

	// Home route
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, `
            <!DOCTYPE html>
            <html>
            <head>
                <title>CRM API Overview</title>
            </head>
            <body>
                <h1>CRM API Overview</h1>
                <p>Welcome to the CRM API. Below are the available endpoints:</p>
                <ul>
                    <li>GET /customers - Retrieve all customers</li>
                    <li>POST /customers - Create a new customer</li>
                    <li>GET /customers/{id} - Retrieve a single customer by ID</li>
                    <li>PUT /customers/{id} - Update customer information</li>
                    <li>DELETE /customers/{id} - Delete a customer by ID</li>
                </ul>
            </body>
            </html>
        `)
	})

	// API routes
	router.HandleFunc("/customers", handler.GetCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", handler.GetCustomerByID).Methods("GET")
	router.HandleFunc("/customers", handler.CreateCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", handler.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", handler.DeleteCustomer).Methods("DELETE")

	return router
}
