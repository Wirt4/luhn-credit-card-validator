package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"main.go/packages/credit_card"
	"main.go/packages/interfaces"
)

type Handler struct {
	validator interfaces.Validator
}

type Payload struct {
	CreditCardNumber string
}

func NewHandler(validator interfaces.Validator) *Handler {
	return &Handler{
		validator: validator,
	}
}

func (h *Handler) HandleGetRequest(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed", http.StatusMethodNotAllowed)
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

	card := credit_card.NewCreditCard()
	card.SetSequence(p.CreditCardNumber)
	response := response{ValidCreditCardNumber: h.validator.IsValid(&card)}
	w.WriteHeader(http.StatusOK)
	encoder.Encode(response)
}

type response struct {
	ValidCreditCardNumber bool
}
