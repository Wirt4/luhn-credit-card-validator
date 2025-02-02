package routing_number

import "testing"

func TestHasInCorrectLength(t *testing.T) {
	number := &RoutingNumber{}
	number.SetSequence("1234-5678-90")
	if number.HasCorrectLength() {
		t.Errorf("Expected false, got true")
	}
}

func TestHasCorrectLength(t *testing.T) {
	number := &RoutingNumber{}
	number.SetSequence("1234-5678-9")
	if !number.HasCorrectLength() {
		t.Errorf("Expected true, got false")
	}
}
