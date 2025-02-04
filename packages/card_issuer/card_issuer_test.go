package card_issuer

import "testing"

func TestVisa(t *testing.T) {
	visa := NewCardIssuer([]int{4}, true)
	if visa.Issuer != "VISA" {
		t.Errorf("Expected VISA, got %v", visa.Issuer)
	}
	if visa.Min != 16 && visa.Max != 16 {
		t.Errorf("Expected 16, got %v", visa.Min)
	}
}

func TestAMX(t *testing.T) {
	amx := NewCardIssuer([]int{3, 7}, true)
	if amx.Issuer != "American Express" {
		t.Errorf("Expected AMX, got %v", amx.Issuer)
	}
	if amx.Min != 15 && amx.Max != 15 {
		t.Errorf("Expected 15, got %v", amx.Min)
	}
}

func TestDinersClub(t *testing.T) {
	dinersClub := NewCardIssuer([]int{5, 5}, true)
	if dinersClub.Issuer != "Mastercard: Diners Club U.S. and Canada" {
		t.Errorf("Expected Diners Club, got %v", dinersClub.Issuer)
	}
	if dinersClub.Min != 16 && dinersClub.Max != 16 {
		t.Errorf("Expected 16, got %v", dinersClub.Min)
	}
}

func TestDancort(t *testing.T) {
	dancort := NewCardIssuer([]int{5, 0, 1, 9, 8}, false)
	if dancort.Issuer != "Dankort" {
		t.Errorf("Expected Dancort, got %v", dancort.Issuer)
	}
	if dancort.Min != 16 && dancort.Max != 16 {
		t.Errorf("Expected 16, got %v", dancort.Min)
	}
}
