package main

import (
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)

	log.Println("ContextVault Go service running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
