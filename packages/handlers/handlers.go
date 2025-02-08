package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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
	card := h.factory.Create()
	is_valid, error := h.isValid(errorHandler.GetParsed(), card)
	if error != nil {
		http.Error(w, error.Error(), http.StatusFailedDependency)
		return
	}
	i := []string{}
	issuers := card.Issuers()
	for _, issuer := range issuers {
		fmt.Println(issuer.Issuer)
		i = append(i, issuer.Issuer)
	}
	response := types.CreditCardResponse{Valid: is_valid, Issuer: strings.Join(i, " ")}
	writeResponse(w, response)
}

func (h *GetHandler) isValid(payload types.CreditCardRequest, card interfaces.DigitSequence) (bool, error) {
	fmt.Print("Card: ", card)
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
