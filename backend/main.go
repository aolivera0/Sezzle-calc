package main

import (
	"log"
	"net/http"
	"sezzle-calc/backend/internal/httpapi"
)

func main() {
	http.HandleFunc("/api/calculate", httpapi.CalculateHandler)

	log.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
