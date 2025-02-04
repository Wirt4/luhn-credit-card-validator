package card_issuer

import (
	"fmt"
	"testing"
)

func TestVisa(t *testing.T) {
	visa := GetCardIssuers([]int{4})
	if visa[0].Issuer != "VISA" {
		t.Errorf("Expected VISA, got %v", visa[0].Issuer)
	}
	if visa[0].Min != 16 && visa[0].Max != 16 {
		t.Errorf("Expected 16, got %v", visa[0].Min)
	}
}

func TestVisaElectron(t *testing.T) {
	visaElectron := GetCardIssuers([]int{4, 1, 7, 5, 0, 0})
	fmt.Println(visaElectron)
	if visaElectron[1].Issuer != "Visa Electron" {
		t.Errorf("Expected Visa Electron, got %v", visaElectron[1].Issuer)
	}
	if visaElectron[1].Min != 16 && visaElectron[1].Max != 16 {
		t.Errorf("Expected 16, got %v", visaElectron[1].Min)
	}
}

func TestAMX(t *testing.T) {
	amx := GetCardIssuers([]int{3, 7})
	if amx[0].Issuer != "American Express" {
		t.Errorf("Expected AMX, got %v", amx[0].Issuer)
	}
	if amx[0].Min != 15 && amx[0].Max != 15 {
		t.Errorf("Expected 15, got %v", amx[0].Min)
	}
}

func TestDinersClub(t *testing.T) {
	dinersClub := GetCardIssuers([]int{5, 5})
	if dinersClub[0].Issuer != "Mastercard: Diners Club U.S. and Canada" {
		t.Errorf("Expected Diners Club, got %v", dinersClub[0].Issuer)
	}
	if dinersClub[0].Min != 16 && dinersClub[0].Max != 16 {
		t.Errorf("Expected 16, got %v", dinersClub[0].Min)
	}
}

func TestDancort(t *testing.T) {
	dancort := GetCardIssuers([]int{5, 0, 1, 9, 8})
	if dancort[0].Issuer != "Dankort" {
		t.Errorf("Expected Dankort, got %v", dancort[0].Issuer)
	}
	if dancort[0].Min != 16 && dancort[0].Max != 16 {
		t.Errorf("Expected 16, got %v", dancort[0].Min)
	}
}
