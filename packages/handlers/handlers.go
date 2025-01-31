package handlers

import (
	"encoding/json"
	"net/http"

	"main.go/packages/credit_card"
	"main.go/packages/interfaces"
)

type Handler struct {
	validator interfaces.Validator
}

func NewHandler(validator interfaces.Validator) *Handler {
	return &Handler{validator: validator}
}

func (h *Handler) HandleGetRequest(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		encoder.Encode(error{ErrorMessage: "Only GET requests are allowed"})
		return
	}

	credit_card := credit_card.NewCreditCard()
	credit_card.SetSequence("1234 5678 9012 3456")

	response := response{ValidCreditCardNumber: h.validator.IsValid(&credit_card)}
	w.WriteHeader(http.StatusOK)
	encoder.Encode(response)
}

type response struct {
	ValidCreditCardNumber bool
}

type error struct {
	ErrorMessage string
}
