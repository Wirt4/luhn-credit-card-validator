package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"main.go/packages/interfaces"
	"main.go/packages/types"
)

type mockCreditCard struct {
	sequence []int
}

type mockFactory struct{}

func (f *mockFactory) NewCreditCard() interfaces.DigitSequence {
	return &mockCreditCard{}
}

func (c *mockCreditCard) SetSequence(sequence string) error {
	for _, v := range sequence {
		if v >= '0' && v <= '9' {
			c.sequence = append(c.sequence, int(v-'0'))
		}
	}
	return nil
}

func (c *mockCreditCard) GetSequence() []int {
	return c.sequence
}

func (c *mockCreditCard) HasCorrectLength() bool {
	return len(c.sequence) == 16
}

type mockValidator struct {
	calledWith interfaces.DigitSequence
	valid      bool
}

func (m *mockValidator) IsValid(sequence interfaces.DigitSequence) (bool, error) {
	m.calledWith = sequence
	return m.valid, nil
}

func TestOnlyShouldAllowGetRequests(t *testing.T) {
	r := constructRequest("POST", types.CreditCardRequest{CreditCardNumber: "1234 5678 9012 3456"})
	w := httptest.NewRecorder()
	expectedBody := "Only GET requests are allowed\n"
	mockHandler := NewHandler(&mockValidator{}, &mockFactory{})

	mockHandler.HandleRequest(w, r)
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
	r := constructRequest("GET", types.CreditCardRequest{CreditCardNumber: "1234 5678 9012 3456"})
	w := httptest.NewRecorder()
	nonExpectedBody := "{\"ErrorMessage\":\"Only GET requests are allowed\"}"
	mockHandler := NewHandler(&mockValidator{valid: true}, &mockFactory{})

	mockHandler.HandleRequest(w, r)
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
	r := constructRequest("GET", types.CreditCardRequest{CreditCardNumber: "1234 5678 9012 3456"})
	w := httptest.NewRecorder()
	expectedBody := "{\"Issuer\":\"\",\"Valid\":true}\n"
	mockHandler := NewHandler(&mockValidator{valid: true}, &mockFactory{})

	mockHandler.HandleRequest(w, r)
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
	r := constructRequest("GET", types.CreditCardRequest{CreditCardNumber: "1234 5678 9012 3456"})
	w := httptest.NewRecorder()
	expectedBody := "{\"Issuer\":\"\",\"Valid\":false}\n"
	mockHandler := NewHandler(&mockValidator{valid: false}, &mockFactory{})

	mockHandler.HandleRequest(w, r)
	response := w.Result()
	body := w.Body.String()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, response.StatusCode)
	}
	if body != expectedBody {
		t.Errorf("Expected body %v, got %v", expectedBody, body)
	}
}

/*
	func TestParametersPassedToHandler(t *testing.T) {
		inputSequence := "4321 8756 2109 6543"
		expected := factories.CreditCardFactory()
		expected.SetSequence(inputSequence)
		r := constructRequest("GET", types.CreditCardRequest{CreditCardNumber: inputSequence})
		w := httptest.NewRecorder()
		mockValidator := &mockValidator{}
		mockHandler := NewHandler(mockValidator)

		mockHandler.HandleRequest(w, r)

		if !reflect.DeepEqual(expected, mockValidator.calledWith) {
			t.Errorf("Expected sequence %v, got %v", expected, mockValidator.calledWith)
		}
	}
*/
func TestHandleGetRequestInvalidMethod(t *testing.T) {
	r := constructRequest("POST", types.CreditCardRequest{CreditCardNumber: "1234 5678 9012 3456"})
	w := httptest.NewRecorder()
	expectedBody := "Only GET requests are allowed\n"
	mockHandler := NewHandler(&mockValidator{}, &mockFactory{})

	mockHandler.HandleRequest(w, r)
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
	r := constructRequest("GET", types.CreditCardRequest{CreditCardNumber: "1234 5678 9012 3456"})
	w := httptest.NewRecorder()
	expectedBody := "{\"Issuer\":\"\",\"Valid\":true}\n"
	mockHandler := NewHandler(&mockValidator{valid: true}, &mockFactory{})

	mockHandler.HandleRequest(w, r)
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
	r := constructRequest("GET", types.CreditCardRequest{CreditCardNumber: "1234 5678 9012 3456"})
	w := httptest.NewRecorder()
	expectedBody := "{\"Issuer\":\"\",\"Valid\":false}\n"
	mockHandler := NewHandler(&mockValidator{valid: false}, &mockFactory{})

	mockHandler.HandleRequest(w, r)
	response := w.Result()
	body := w.Body.String()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
	}
	if body != expectedBody {
		t.Errorf("Expected body %v, got %v", expectedBody, body)
	}
}

func constructRequest(method string, body types.CreditCardRequest) *http.Request {

	out, _ := json.Marshal(body)

	req, _ := http.NewRequest(method, "http://example.com", bytes.NewBuffer(out))
	return req
}
