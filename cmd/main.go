package main

import (
	"crm-app/internal/models"
	"crm-app/internal/routes"
	"log"
	"net/http"
)

var customers = []models.Customer{
	{ID: 1, Name: "John Doe", Email: "john.doe@example.com"},
	{ID: 2, Name: "Jane Smith", Email: "jane.smith@example.com"},
}

func main() {
	router := routes.RegisterRoutes()
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
