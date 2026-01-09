package handlers

import (
	"net/http"

	"contextvault/go/internal/clients"
)

func PythonHealthHandler(pythonClient *clients.PythonClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := pythonClient.Health(); err != nil {
			http.Error(w, "python service unreachable", http.StatusServiceUnavailable)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("python ok"))
	}
}
