package handlers

import (
	"encoding/json"
	"io"
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

	var p Payload

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &p); err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	h.credit_card_container.SetSequence(p.CreditCardNumber)

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
