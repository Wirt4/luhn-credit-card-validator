package number_parser

import (
	"reflect"
	"testing"

	"main.go/packages/types"
)

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
		Name: "American Express",
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
