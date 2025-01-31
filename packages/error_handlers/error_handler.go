package error_handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"main.go/packages/types"
)

type GetErrorHandler struct {
	Message  string
	Code     int
	HasError bool
	Parsed   types.CreditCardPayload
}

func NewErrorHandler() *GetErrorHandler {
	return &GetErrorHandler{
		Message:  "",
		Code:     http.StatusOK,
		HasError: false,
		Parsed:   types.CreditCardPayload{},
	}
}

func (h *GetErrorHandler) CheckMethod(method string) {
	if method != http.MethodGet && !h.HasError {
		h.setError("Only GET requests are allowed", http.StatusMethodNotAllowed)
	}
}

func (h *GetErrorHandler) CheckBody(body io.ReadCloser) {
	if h.HasError {
		return
	}

	b, err := io.ReadAll(body)
	if err != nil {
		h.setError("Error reading request body", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(b, &h.Parsed); err != nil {
		h.setError("Error parsing request body", http.StatusBadRequest)
	}
}

func (h *GetErrorHandler) setError(message string, code int) {
	h.Message = message
	h.Code = code
	h.HasError = true
}
