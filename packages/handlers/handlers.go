package handlers

import "net/http"

func HandleGetRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Only GET requests are allowed\n"))
		return
	}
	w.WriteHeader(http.StatusOK)
}
