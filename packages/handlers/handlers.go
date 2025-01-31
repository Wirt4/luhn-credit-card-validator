package handlers

import "net/http"

func HandleGetRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Only GET requests are allowed\n"))
}
