package handlers

import (
	"encoding/json"
	"net/http"

	"main.go/packages/interfaces"
)

type Handler struct {
	validator             interfaces.Validator
	credit_card_container interfaces.DigitSequence
}

type Payload struct {
	CreditCardNumber string
}

func NewHandler(validator interfaces.Validator, container interfaces.DigitSequence) *Handler {
	return &Handler{
		validator:             validator,
		credit_card_container: container,
	}
}

func (h *Handler) HandleGetRequest(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		encoder.Encode(error{ErrorMessage: "Only GET requests are allowed"})
		return
	}

	h.credit_card_container.SetSequence("1234 5678 9012 3456")

	response := response{ValidCreditCardNumber: h.validator.IsValid(h.credit_card_container)}
	w.WriteHeader(http.StatusOK)
	encoder.Encode(response)
}

type response struct {
	ValidCreditCardNumber bool
}

type error struct {
	ErrorMessage string
}
