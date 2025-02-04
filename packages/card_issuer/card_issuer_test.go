package card_issuer

import (
	"testing"

	"main.go/packages/types"
)

type mockVisor struct {
	visited []types.CardIssuer
}

func (v *mockVisor) Traverse(sequence []int, root *types.Node) {}
func (v *mockVisor) GetVisited() []types.CardIssuer {
	return v.visited
}
func TestVisa(t *testing.T) {
	visitor := &mockVisor{
		visited: []types.CardIssuer{
			{Issuer: "VISA", Min: 16, Max: 16}, {Issuer: "Visa Electron", Min: 16, Max: 16},
		},
	}
	visa := GetCardIssuers([]int{4}, visitor)
	if visa[0].Issuer != "VISA" {
		t.Errorf("Expected VISA, got %v", visa[0].Issuer)
	}
	if visa[1].Issuer != "Visa Electron" {
		t.Errorf("Expected Visa Electron, got %v", visa[1].Issuer)
	}
	if visa[0].Min != 16 && visa[0].Max != 16 {
		t.Errorf("Expected 16, got %v", visa[0].Min)
	}
}

func TestAMX(t *testing.T) {
	visitor := &mockVisor{
		visited: []types.CardIssuer{
			{Issuer: "American Express", Min: 15, Max: 15},
		},
	}
	amx := GetCardIssuers([]int{3, 7}, visitor)
	if amx[0].Issuer != "American Express" {
		t.Errorf("Expected AMX, got %v", amx[0].Issuer)
	}
	if amx[0].Min != 15 && amx[0].Max != 15 {
		t.Errorf("Expected 15, got %v", amx[0].Min)
	}
}
