package credit_card

import (
	"reflect"
	"testing"
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
	card := &CreditCard{}
	card.SetSequence("1234 5678 9111 1111 90")
	if card.HasCorrectLength() {
		t.Errorf("Expected false, got true")
	}
}

func TestHasCorrectLengthWithZero(t *testing.T) {
	card := &CreditCard{}
	card.SetSequence("1234 5678 9111 1101")
	if !card.HasCorrectLength() {
		t.Errorf("Expected true, got false")
	}
}
