package card_issuer

import (
	"reflect"
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
	visa := visitor.GetVisited()
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

func TestParseEntryVisa(t *testing.T) {
	expected := types.ProviderData{
		Name:              "VISA",
		IINs:              []int{4},
		MaxSequenceLength: 16,
		MinSequenceLength: 16,
	}
	actual := ParseEntry("VISA 4 16")
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestParseEntryDinersClub(t *testing.T) {
	expected := types.ProviderData{
		Name:              "Mastercard: Diners Club U.S. and Canada",
		IINs:              []int{55},
		MaxSequenceLength: 16,
		MinSequenceLength: 16,
	}
	actual := ParseEntry("Mastercard: Diners Club U.S. and Canada 55 16")
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestParsedEntryMultipleIINs(t *testing.T) {
	expected := types.ProviderData{
		Name:              "Mastercard",
		IINs:              []int{2221, 2222, 2223},
		MaxSequenceLength: 16,
		MinSequenceLength: 16,
	}
	actual := ParseEntry("Mastercard 2221-2223 16")
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestParseAMX(t *testing.T) {
	expected := types.ProviderData{
		Name:              "American Express",
		IINs:              []int{34, 35},
		MaxSequenceLength: 15,
		MinSequenceLength: 15,
	}
	actual := ParseEntry("American Express [34, 35] 15")
	if !reflect.DeepEqual(expected.Name, actual.Name) {
		t.Errorf("Expected %v, got %v", expected.Name, actual.Name)
	}
}

func TestSplit(t *testing.T) {
	expected := []string{"American Express", "[34, 35]", "15"}
	actual := split("American Express [34, 35] 15")
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

}
