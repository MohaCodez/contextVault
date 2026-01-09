package main

import (
	"log"
	"net/http"
	"os"

	"contextvault/go/internal/api/handlers"
	"contextvault/go/internal/clients"
)

func main() {
	mux := http.NewServeMux()

	pythonURL := os.Getenv("PYTHON_SERVICE_URL")
	if pythonURL == "" {
		pythonURL = "http://localhost:8000"
	}

	pythonClient := clients.NewPythonClient(pythonURL)

	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/python-health", handlers.PythonHealthHandler(pythonClient))

	log.Println("ContextVault Go service running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
