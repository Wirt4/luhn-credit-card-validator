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
	factory   interfaces.DigitSequenceFactory
}

func NewHandler(validator interfaces.Validator, factory interfaces.DigitSequenceFactory) *GetHandler {
	return &GetHandler{
		validator: validator,
		factory:   factory,
	}
}

func (h *GetHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	errHandlerFactory := &factories.ErrorHandlerFactory{}
	errorHandler := errHandlerFactory.Create()
	errorHandler.CheckMethod(r.Method)
	errorHandler.CheckBody(r.Body)
	if errorHandler.HasError() {
		message := errorHandler.GetMessage()
		http.Error(w, message, http.StatusMethodNotAllowed)
		return
	}
	is_valid, error := h.isValid(errorHandler.GetParsed())
	if error != nil {
		http.Error(w, error.Error(), http.StatusFailedDependency)
		return
	}
	response := types.CreditCardResponse{Valid: is_valid}
	writeResponse(w, response)
}

func (h *GetHandler) isValid(payload types.CreditCardRequest) (bool, error) {
	card := h.factory.NewCreditCard()
	error := card.SetSequence(payload.CreditCardNumber)
	if error != nil {
		return false, error
	}
	return h.validator.IsValid(card)
}
func writeResponse(w http.ResponseWriter, response types.CreditCardResponse) {
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
