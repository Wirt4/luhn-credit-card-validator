package routing_checksum

import (
	"testing"

	"main.go/packages/routing_number"
)

func TestIsInValidLength(t *testing.T) {
	input := &routing_number.RoutingNumber{}
	input.SetSequence("1234")
	routing_checksum := &RoutingChecksum{}

	if routing_checksum.IsValid(input) {
		t.Errorf("Expected false, got true")
	}
}

func TestIsInValid(t *testing.T) {
	input := &routing_number.RoutingNumber{}
	input.SetSequence("111000024")
	routing_checksum := &RoutingChecksum{}

	if routing_checksum.IsValid(input) {
		t.Errorf("Expected false, got true")
	}
}
