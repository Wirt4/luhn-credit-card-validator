package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"main.go/packages/interfaces"
)

type MockDigitSequence struct {
	sequence string
}

func (m *MockDigitSequence) SetSequence(sequence string) {
	m.sequence = sequence
}

func (m *MockDigitSequence) GetSequence() []int {
	return []int{}
}

func (m *MockDigitSequence) HasCorrectLength() bool {
	return true
}

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
	expectedBody := "{\"ErrorMessage\":\"Only GET requests are allowed\"}\n"
	mockHandler := NewHandler(&mockValidator{}, &MockDigitSequence{})

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
	mockHandler := NewHandler(&mockValidator{valid: true}, &MockDigitSequence{})

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
	mockHandler := NewHandler(&mockValidator{valid: true}, &MockDigitSequence{})

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
	mockHandler := NewHandler(&mockValidator{valid: false}, &MockDigitSequence{})

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
	expectedSequence := "4321 8756 2109 6543"
	r := constructRequest("GET", Payload{CreditCardNumber: expectedSequence})
	w := httptest.NewRecorder()
	mockValidator := &mockValidator{}
	mockHandler := NewHandler(mockValidator, &MockDigitSequence{})

	mockHandler.HandleGetRequest(w, r)

	if mockValidator.calledWith.(*MockDigitSequence).sequence != expectedSequence {
		t.Errorf("Expected sequence %v, got %v", expectedSequence, mockValidator.calledWith.(*MockDigitSequence).sequence)
	}
}

func constructRequest(method string, body Payload) *http.Request {

	out, _ := json.Marshal(body)

	req, _ := http.NewRequest(method, "http://example.com", bytes.NewBuffer(out))
	return req
}
