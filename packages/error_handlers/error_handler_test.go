package error_handlers

import (
	"bytes"
	"io"
	"net/http"
	"reflect"
	"testing"

	"main.go/packages/types"
)

func TestCheckMethod(t *testing.T) {
	tests := []struct {
		method          string
		expectedError   bool
		expectedMessage string
		expectedCode    int
		initialHasError bool
	}{
		{http.MethodGet, false, "", http.StatusOK, false},
		{http.MethodPost, true, "Only GET requests are allowed", http.StatusMethodNotAllowed, false},
		{http.MethodPut, true, "Only GET requests are allowed", http.StatusMethodNotAllowed, false},
		{http.MethodDelete, true, "Only GET requests are allowed", http.StatusMethodNotAllowed, false},
		{http.MethodGet, true, "", http.StatusOK, true},
	}
	for _, test := range tests {
		t.Run(test.method, func(t *testing.T) {
			errorHandler := &GetErrorHandler{
				message:  "",
				code:     http.StatusOK,
				hasError: test.initialHasError,
			}
			errorHandler.CheckMethod(test.method)

			if errorHandler.HasError() != test.expectedError {
				t.Errorf("Expected HasError to be %v, got %v", test.expectedError, errorHandler.HasError())
			}
			if errorHandler.GetMessage() != test.expectedMessage {
				t.Errorf("Expected Message to be %v, got %v", test.expectedMessage, errorHandler.GetMessage())
			}
			if errorHandler.GetCode() != test.expectedCode {
				t.Errorf("Expected Code to be %v, got %v", test.expectedCode, errorHandler.GetCode())
			}
		})
	}
}

func TestCheckBody(t *testing.T) {
	tests := []struct {
		name            string
		body            io.ReadCloser
		initialHasError bool
		expectedError   bool
		expectedMessage string
		expectedCode    int
		expectedParsed  types.CreditCardRequest
	}{
		{
			name:            "HasError already true",
			body:            io.NopCloser(bytes.NewReader([]byte(`{"CreditCardNumber": "1234 5678 9012 3456"}`))),
			initialHasError: true,
			expectedError:   true,
			expectedMessage: "",
			expectedCode:    http.StatusOK,
			expectedParsed:  types.CreditCardRequest{},
		},
		{
			name:            "Error reading body",
			body:            io.NopCloser(&errorReader{}),
			initialHasError: false,
			expectedError:   true,
			expectedMessage: "Error reading request body",
			expectedCode:    http.StatusInternalServerError,
			expectedParsed:  types.CreditCardRequest{},
		},
		{
			name:            "Error parsing body",
			body:            io.NopCloser(bytes.NewReader([]byte(`invalid json`))),
			initialHasError: false,
			expectedError:   true,
			expectedMessage: "Error parsing request body",
			expectedCode:    http.StatusBadRequest,
			expectedParsed:  types.CreditCardRequest{},
		},
		{
			name:            "Successful parsing",
			body:            io.NopCloser(bytes.NewReader([]byte(`{"CreditCardNumber": "1234 5678 9012 3456"}`))),
			initialHasError: false,
			expectedError:   false,
			expectedMessage: "",
			expectedCode:    http.StatusOK,
			expectedParsed:  types.CreditCardRequest{CreditCardNumber: "1234 5678 9012 3456"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			errorHandler := &GetErrorHandler{
				message:  "",
				code:     http.StatusOK,
				hasError: test.initialHasError,
			}

			errorHandler.CheckBody(test.body)

			if errorHandler.HasError() != test.expectedError {
				t.Errorf("Expected HasError to be %v, got %v", test.expectedError, errorHandler.HasError())
			}
			if errorHandler.GetMessage() != test.expectedMessage {
				t.Errorf("Expected Message to be %v, got %v", test.expectedMessage, errorHandler.GetMessage())
			}
			if errorHandler.GetCode() != test.expectedCode {
				t.Errorf("Expected Code to be %v, got %v", test.expectedCode, errorHandler.GetCode())
			}
			if !reflect.DeepEqual(errorHandler.GetParsed(), test.expectedParsed) {
				t.Errorf("Expected Parsed to be %v, got %v", test.expectedParsed, errorHandler.GetParsed())
			}
		})
	}
}

// errorReader is a helper type that always returns an error when Read is called.
type errorReader struct{}

func (e *errorReader) Read(p []byte) (n int, err error) {
	return 0, io.ErrUnexpectedEOF
}
