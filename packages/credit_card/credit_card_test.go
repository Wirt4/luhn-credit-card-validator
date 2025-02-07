package credit_card

import (
	"reflect"
	"testing"

	"main.go/packages/types"
)

func TestUndelimitedStringInput(t *testing.T) {
	expected := []int{4, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	card := &CreditCard{}

	card.SetSequence("4111111111111111")
	actual := card.GetSequence()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestUndelimitedStringInputDifferentData(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 1, 1, 1, 1, 1, 1}
	card := &CreditCard{}

	card.SetSequence("1234567891111111")
	actual := card.GetSequence()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestDelimitedStringInput(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 1, 1, 1, 1, 1, 1}
	card := &CreditCard{}
	card.SetSequence("1234 5678 9111 1111")
	actual := card.GetSequence()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestHasCorrectLength(t *testing.T) {
	card := &CreditCard{
		issuers: []types.CardIssuer{{
			Min: 16,
			Max: 16,
		}},
	}
	card.SetSequence("1234 5678 9111 1111 90")
	if card.HasCorrectLength() {
		t.Errorf("Expected false, got true")
	}
}

func TestHasCorrectLengthAMX(t *testing.T) {
	card := &CreditCard{
		issuers: []types.CardIssuer{{
			Min: 15,
			Max: 15,
		}},
	}
	card.SetSequence("3434 5678 9111 1111")
	if card.HasCorrectLength() {
		t.Errorf("Expected false, got true")
	}
}

func TestHasCorrectLengthWithZero(t *testing.T) {
	card := &CreditCard{}
	card.SetSequence("4234 5678 9111 1101")
	if !card.HasCorrectLength() {
		t.Errorf("Expected true, got false")
	}
}

func TestCorrectLengthWithMultipleHits(t *testing.T) {
	card := &CreditCard{}
	//Integrated with AMX code
	card.SetSequence("3456 5678 9111 198")
	if !card.HasCorrectLength() {
		t.Errorf("Expected true, got false")
	}
	if len(card.issuers) != 1 {
		t.Errorf("Expected 1, got %d", len(card.issuers))
	}
}

func TestMiddleofRange(t *testing.T) {
	card := &CreditCard{}
	//Integrated with China Union Pay code
	card.SetSequence("6256 5678 9111 1989 8")
	if !card.HasCorrectLength() {
		t.Errorf("Expected true, got false")
	}
}
