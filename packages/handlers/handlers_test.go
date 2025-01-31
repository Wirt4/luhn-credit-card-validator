package handlers

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"main.go/packages/interfaces"
)

type mockValidator struct {
	valid bool
}

func (m *mockValidator) IsValid(sequence interfaces.DigitSequence) bool {
	return m.valid
}

func TestOnlyShouldAllowGetRequests(t *testing.T) {
	r := http.Request{
		Method: "POST",
	}
	w := httptest.NewRecorder()
	expectedBody := "{\"ErrorMessage\":\"Only GET requests are allowed\"}\n"
	mockHandler := Handler{
		validator: &mockValidator{},
	}

	mockHandler.HandleGetRequest(w, &r)
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
	r := http.Request{
		Method: "GET",
	}
	w := httptest.NewRecorder()
	nonExpectedBody := "{\"ErrorMessage\":\"Only GET requests are allowed\"}"

	mockHandler := Handler{
		validator: &mockValidator{
			valid: true,
		},
	}

	mockHandler.HandleGetRequest(w, &r)
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
	r := http.Request{
		Method: "GET",
	}
	w := httptest.NewRecorder()
	expectedBody := "{\"ValidCreditCardNumber\":true}\n"

	mockHandler := Handler{
		validator: &mockValidator{
			valid: true,
		},
	}

	mockHandler.HandleGetRequest(w, &r)
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
	r := http.Request{
		Method: "GET",
	}
	w := httptest.NewRecorder()
	expectedBody := "{\"ValidCreditCardNumber\":false}\n"

	mockHandler := Handler{
		validator: &mockValidator{
			valid: false,
		},
	}

	mockHandler.HandleGetRequest(w, &r)
	response := w.Result()
	body := w.Body.String()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, response.StatusCode)
	}
	if body != expectedBody {
		t.Errorf("Expected body %v, got %v", expectedBody, body)
	}
}
