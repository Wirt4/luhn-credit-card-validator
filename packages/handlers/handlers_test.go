package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"main.go/packages/credit_card"
	"main.go/packages/interfaces"
)

type mockValidator struct {
	calledWith interfaces.DigitSequence
	valid      bool
}

func (m *mockValidator) IsValid(sequence interfaces.DigitSequence) bool {
	m.calledWith = sequence
	return m.valid
}

func TestOnlyShouldAllowGetRequests(t *testing.T) {
	r := constructRequest("POST", Payload{CreditCardNumber: "1234 5678 9012 3456"})
	w := httptest.NewRecorder()
	expectedBody := "Only GET requests are allowed\n"
	mockHandler := NewHandler(&mockValidator{})

	mockHandler.HandleGetRequest(w, r)
	response := w.Result()
	body := w.Body.String()

	if response.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, response.StatusCode)
	}
	if !reflect.DeepEqual(body, expectedBody) {
		t.Errorf("Expected body %v, got %v", expectedBody, body)
	}
}

func TestOnlyShouldAllowGetRequestsDifferentData(t *testing.T) {
	r := constructRequest("GET", Payload{CreditCardNumber: "1234 5678 9012 3456"})
	w := httptest.NewRecorder()
	nonExpectedBody := "{\"ErrorMessage\":\"Only GET requests are allowed\"}"
	mockHandler := NewHandler(&mockValidator{valid: true})

	mockHandler.HandleGetRequest(w, r)
	response := w.Result()
	body := w.Body.String()

	if response.StatusCode == http.StatusMethodNotAllowed {
		t.Errorf("Expected status code other than %v", response.StatusCode)
	}
	if body == nonExpectedBody {
		t.Errorf("Expected body other than \"%v\"", nonExpectedBody)
	}
}

func TestIfValidatorReturnsTrue(t *testing.T) {
	r := constructRequest("GET", Payload{CreditCardNumber: "1234 5678 9012 3456"})
	w := httptest.NewRecorder()
	expectedBody := "{\"ValidCreditCardNumber\":true}\n"
	mockHandler := NewHandler(&mockValidator{valid: true})

	mockHandler.HandleGetRequest(w, r)
	response := w.Result()
	body := w.Body.String()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, response.StatusCode)
	}
	if body != expectedBody {
		t.Errorf("Expected body %v, got %v", expectedBody, body)
	}
}

func TestIfValidatorReturnsFalse(t *testing.T) {
	r := constructRequest("GET", Payload{CreditCardNumber: "1234 5678 9012 3456"})
	w := httptest.NewRecorder()
	expectedBody := "{\"ValidCreditCardNumber\":false}\n"
	mockHandler := NewHandler(&mockValidator{valid: false})

	mockHandler.HandleGetRequest(w, r)
	response := w.Result()
	body := w.Body.String()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, response.StatusCode)
	}
	if body != expectedBody {
		t.Errorf("Expected body %v, got %v", expectedBody, body)
	}
}

func TestParametersPassedToHandler(t *testing.T) {
	inputSequence := "4321 8756 2109 6543"
	expected := credit_card.NewCreditCard()
	expected.SetSequence(inputSequence)
	r := constructRequest("GET", Payload{CreditCardNumber: inputSequence})
	w := httptest.NewRecorder()
	mockValidator := &mockValidator{}
	mockHandler := NewHandler(mockValidator)

	mockHandler.HandleGetRequest(w, r)

	if !reflect.DeepEqual(&expected, mockValidator.calledWith) {
		t.Errorf("Expected sequence %v, got %v", expected, mockValidator.calledWith)
	}
}
func TestHandleGetRequestInvalidMethod(t *testing.T) {
	r := constructRequest("POST", Payload{CreditCardNumber: "1234 5678 9012 3456"})
	w := httptest.NewRecorder()
	expectedBody := "Only GET requests are allowed\n"
	mockHandler := NewHandler(&mockValidator{})

	mockHandler.HandleGetRequest(w, r)
	response := w.Result()
	body := w.Body.String()

	if response.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, response.StatusCode)
	}
	if body != expectedBody {
		t.Errorf("Expected body %v, got %v", expectedBody, body)
	}
}

func TestHandleGetRequestValidCreditCard(t *testing.T) {
	r := constructRequest("GET", Payload{CreditCardNumber: "1234 5678 9012 3456"})
	w := httptest.NewRecorder()
	expectedBody := "{\"ValidCreditCardNumber\":true}\n"
	mockHandler := NewHandler(&mockValidator{valid: true})

	mockHandler.HandleGetRequest(w, r)
	response := w.Result()
	body := w.Body.String()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
	}
	if body != expectedBody {
		t.Errorf("Expected body %v, got %v", expectedBody, body)
	}
}

func TestHandleGetRequestInvalidCreditCard(t *testing.T) {
	r := constructRequest("GET", Payload{CreditCardNumber: "1234 5678 9012 3456"})
	w := httptest.NewRecorder()
	expectedBody := "{\"ValidCreditCardNumber\":false}\n"
	mockHandler := NewHandler(&mockValidator{valid: false})

	mockHandler.HandleGetRequest(w, r)
	response := w.Result()
	body := w.Body.String()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
	}
	if body != expectedBody {
		t.Errorf("Expected body %v, got %v", expectedBody, body)
	}
}

func constructRequest(method string, body Payload) *http.Request {

	out, _ := json.Marshal(body)

	req, _ := http.NewRequest(method, "http://example.com", bytes.NewBuffer(out))
	return req
}
