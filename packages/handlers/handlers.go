package handlers

import (
	"encoding/json"
	"net/http"

	"main.go/packages/factories"
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

func (h *GetHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	errorHandler := factories.ErrorHandlerFactory()
	errorHandler.CheckMethod(r.Method)
	errorHandler.CheckBody(r.Body)
	if errorHandler.HasError() {
		message := errorHandler.GetMessage()
		http.Error(w, message, http.StatusMethodNotAllowed)
		return
	}
	is_valid := h.isValid(errorHandler.GetParsed())
	response := response{ValidCreditCardNumber: is_valid}
	writeResponse(w, response)
}

func (h *GetHandler) isValid(payload types.CreditCardRequest) bool {
	card := factories.CreditCardFactory()
	card.SetSequence(payload.CreditCardNumber)
	return h.validator.IsValid(card)
}
func writeResponse(w http.ResponseWriter, response response) {
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
