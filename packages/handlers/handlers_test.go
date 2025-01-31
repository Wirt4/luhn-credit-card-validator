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
	sequence []int
}

func (m MockDigitSequence) GetSequence() []int {
	return m.sequence
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
	mockHandler := Handler{
		validator: &mockValidator{},
	}

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

	mockHandler := Handler{
		validator: &mockValidator{
			valid: true,
		},
	}

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

	mockHandler := Handler{
		validator: &mockValidator{
			valid: false,
		},
	}

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
	r := constructRequest("GET", Payload{CreditCardNumber: "1234 5678 9012 3456"})
	w := httptest.NewRecorder()
	mockValidator := &mockValidator{}
	mockHandler := NewHandler(mockValidator)
	expectedArgument := MockDigitSequence{
		sequence: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6},
	}

	mockHandler.HandleGetRequest(w, r)

	if !reflect.DeepEqual(mockValidator.calledWith.GetSequence(), expectedArgument.GetSequence()) {
		t.Errorf("Expected digits object of  %v, got %v", expectedArgument, mockValidator.calledWith)
	}
}

func constructRequest(method string, body Payload) *http.Request {

	out, _ := json.Marshal(body)

	req, _ := http.NewRequest(method, "http://example.com", bytes.NewBuffer(out))
	return req
}
