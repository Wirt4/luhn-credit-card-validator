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

type error struct {
	ErrorMessage string
}

func (h *Handler) HandleGetRequest(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		encoder.Encode(error{ErrorMessage: "Only GET requests are allowed"})
		return
	}

	w.WriteHeader(http.StatusOK)
	response := response{ValidCreditCardNumber: true}
	encoder.Encode(response)
}
