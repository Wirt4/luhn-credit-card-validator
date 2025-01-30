package credit_card

import (
	"reflect"
	"testing"
)

func TestUndelimitedStringInput(t *testing.T) {
	expected := []int{4, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	card := NewCreditCard()

	card.SetSequence("4111111111111111")
	actual := card.GetSequence()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
