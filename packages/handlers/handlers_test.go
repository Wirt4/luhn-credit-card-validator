package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOnlyShouldAllowGetRequests(t *testing.T) {
	r := http.Request{
		Method: "POST",
	}
	w := httptest.NewRecorder()
	expectedBody := "Only GET requests are allowed\n"

	HandleGetRequest(w, &r)
	response := w.Result()
	body := w.Body.String()

	if response.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, response.StatusCode)
	}
	if body != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, body)
	}
}

func TestOnlyShouldAllowGetRequestsDifferentData(t *testing.T) {
	r := http.Request{
		Method: "GET",
	}
	w := httptest.NewRecorder()
	nonExpectedBody := "Only GET requests are allowed\n"

	HandleGetRequest(w, &r)
	response := w.Result()
	body := w.Body.String()

	if response.StatusCode == http.StatusMethodNotAllowed {
		t.Errorf("Expected status code other than %v", response.StatusCode)
	}
	if body == nonExpectedBody {
		t.Errorf("Expected body other than \"%v\"", nonExpectedBody)
	}
}
