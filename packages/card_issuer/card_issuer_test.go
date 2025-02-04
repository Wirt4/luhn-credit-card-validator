package card_issuer

import "testing"

func TestVisa(t *testing.T) {
	visa := NewCardIssuer([]int{4}, true)
	if visa.Issuer != "VISA" {
		t.Errorf("Expected VISA, got %v", visa.Issuer)
	}
	if visa.SequenceLength != 16 {
		t.Errorf("Expected 16, got %v", visa.SequenceLength)
	}
}

func TestAMX(t *testing.T) {
	amx := NewCardIssuer([]int{3, 7}, true)
	if amx.Issuer != "American Express" {
		t.Errorf("Expected AMX, got %v", amx.Issuer)
	}
	if amx.SequenceLength != 15 {
		t.Errorf("Expected 15, got %v", amx.SequenceLength)
	}
}

func TestDinersClub(t *testing.T) {
	dinersClub := NewCardIssuer([]int{5, 5}, true)
	if dinersClub.Issuer != "Mastercard: Diners Club U.S. and Canada" {
		t.Errorf("Expected Diners Club, got %v", dinersClub.Issuer)
	}
	if dinersClub.SequenceLength != 16 {
		t.Errorf("Expected 16, got %v", dinersClub.SequenceLength)
	}
}

func TestDancort(t *testing.T) {
	dancort := NewCardIssuer([]int{5, 0, 1, 9, 8}, false)
	if dancort.Issuer != "Dankort" {
		t.Errorf("Expected Dancort, got %v", dancort.Issuer)
	}
	if dancort.SequenceLength != 16 {
		t.Errorf("Expected 16, got %v", dancort.SequenceLength)
	}
}
