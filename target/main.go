package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// handler for the GET request
func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters
	statusParam := r.URL.Query().Get("index")

	// Log the incoming request
	log.Printf("Received request: StatusParam: %s", statusParam)

	// Convert query parameter to integer
	statusCode, err := strconv.Atoi(statusParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid status code")
		return
	}

	// Respond based on the status parameter
	switch statusCode {
	case http.StatusOK:
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Status 200: OK")
	case http.StatusNoContent:
		w.WriteHeader(http.StatusNoContent)
	case http.StatusBadRequest:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Status 400: Bad Request")
	case http.StatusInternalServerError:
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Status 500: Internal Server Error")
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Status 404: Not Found")
	}
}

func main() {
	// Define the endpoint and attach the handler
	http.HandleFunc("/api", handleRequest)

	// Start the HTTP server
	port := 8081
	log.Printf("Starting target server on port %d", port)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
