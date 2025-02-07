package luhn

import "testing"

type mockDigitSeqence struct {
	sequence            []int
	correctNumberLength bool
}

func (m *mockDigitSeqence) SetSequence(sequence string) error {
	return nil
}

func (m *mockDigitSeqence) GetSequence() []int {
	return m.sequence
}

func (m *mockDigitSeqence) HasCorrectLength() bool {
	return m.correctNumberLength
}

func TestIsValidSmallCase(t *testing.T) {
	mockNumbers := &mockDigitSeqence{
		sequence:            []int{1, 2, 5},
		correctNumberLength: true,
	}

	validaor := &LuhnValidator{}
	actual, _ := validaor.IsValid(mockNumbers)

	if actual != true {
		t.Errorf("%v should be valid", mockNumbers.GetSequence())
	}
}

func TestIsInvalid(t *testing.T) {
	mockNumbers := &mockDigitSeqence{
		sequence:            []int{1, 2, 4},
		correctNumberLength: true,
	}

	validaor := &LuhnValidator{}
	actual, _ := validaor.IsValid(mockNumbers)

	if actual != false {
		t.Errorf("%v should not be valid", mockNumbers.GetSequence())
	}
}

func TestIsValid(t *testing.T) {
	mockNumbers := &mockDigitSeqence{
		sequence:            []int{1, 7, 8, 9, 3, 7, 2, 9, 9, 7, 4},
		correctNumberLength: true,
	}

	validaor := &LuhnValidator{}
	actual, _ := validaor.IsValid(mockNumbers)

	if actual != true {
		t.Errorf("%v should be valid", mockNumbers.GetSequence())
	}
}

func TestEmptySequence(t *testing.T) {
	mockNumbers := &mockDigitSeqence{
		sequence:            []int{},
		correctNumberLength: true,
	}

	validaor := &LuhnValidator{}
	actual, _ := validaor.IsValid(mockNumbers)

	if actual != false {
		t.Errorf("%v should not be valid", mockNumbers.GetSequence())
	}
}

func TestIsInvalidLength(t *testing.T) {
	mockNumbers := &mockDigitSeqence{
		sequence:            []int{1, 2, 5},
		correctNumberLength: false,
	}

	validaor := &LuhnValidator{}
	actual, _ := validaor.IsValid(mockNumbers)

	if actual != false {
		t.Errorf("%v should not be valid", mockNumbers.GetSequence())
	}
}
