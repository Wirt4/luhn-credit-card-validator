package handlers

import (
	"encoding/json"
	"net/http"

	"main.go/packages/interfaces"
)

type Handler struct {
	validator interfaces.Validator
}

type response struct {
	ValidCreditCardNumber bool
}

func (h *Handler) HandleGetRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Only GET requests are allowed\n"))
		return
	}

	w.WriteHeader(http.StatusOK)
	response := response{ValidCreditCardNumber: true}
	json.NewEncoder(w).Encode(response)
}
