package handlers

import (
	"encoding/json"
	"net/http"

	"main.go/packages/credit_card"
	"main.go/packages/error_handlers"
	"main.go/packages/interfaces"
	"main.go/packages/types"
)

type GetHandler struct {
	validator interfaces.Validator
}

type response struct {
	ValidCreditCardNumber bool
}

func NewHandler(validator interfaces.Validator) *GetHandler {
	return &GetHandler{
		validator: validator,
	}
}

func (h *GetHandler) HandleGetRequest(w http.ResponseWriter, r *http.Request) {
	errorHandler := error_handlers.NewErrorHandler()
	errorHandler.CheckMethod(r.Method)
	errorHandler.CheckBody(r.Body)
	if errorHandler.HasError {
		message := errorHandler.Message
		http.Error(w, message, http.StatusMethodNotAllowed)
		return
	}
	is_valid := h.isValid(errorHandler.Parsed)
	response := response{ValidCreditCardNumber: is_valid}
	writeResponse(w, response)
}

func (h *GetHandler) isValid(payload types.CreditCardPayload) bool {
	card := credit_card.NewCreditCard()
	card.SetSequence(payload.CreditCardNumber)
	return h.validator.IsValid(&card)
}
func writeResponse(w http.ResponseWriter, response response) {
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
