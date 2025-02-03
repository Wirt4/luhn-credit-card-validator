package error_handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"main.go/packages/types"
)

type GetErrorHandler struct {
	message  string
	code     int
	hasError bool
	parsed   types.CreditCardRequest
}

func NewErrorHandler() *GetErrorHandler {
	return &GetErrorHandler{
		message:  "",
		code:     http.StatusOK,
		hasError: false,
		parsed:   types.CreditCardRequest{},
	}
}

func (h *GetErrorHandler) CheckMethod(method string) {
	if method != http.MethodGet && !h.hasError {
		h.setError("Only GET requests are allowed", http.StatusMethodNotAllowed)
	}
}

func (h *GetErrorHandler) CheckBody(body io.ReadCloser) {
	if h.hasError {
		return
	}

	b, err := io.ReadAll(body)
	if err != nil {
		h.setError("Error reading request body", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(b, &h.parsed); err != nil {
		h.setError("Error parsing request body", http.StatusBadRequest)
	}
}

func (h *GetErrorHandler) HasError() bool {
	return h.hasError
}

func (h *GetErrorHandler) GetMessage() string {
	return h.message
}

func (h *GetErrorHandler) GetCode() int {
	return h.code
}

func (h *GetErrorHandler) GetParsed() types.CreditCardRequest {
	return h.parsed
}

func (h *GetErrorHandler) setError(message string, code int) {
	h.message = message
	h.code = code
	h.hasError = true
}
