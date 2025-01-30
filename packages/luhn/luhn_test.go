package luhn

import "testing"

type mockDigitSeqence struct {
	sequence []int
}

func (m *mockDigitSeqence) SetSequence(sequence string) {

}

func (m *mockDigitSeqence) GetSequence() []int {
	return m.sequence
}

func TestIsValid(t *testing.T) {
	mockNumbers := &mockDigitSeqence{
		sequence: []int{1, 2, 5},
	}

	actual := IsValid(mockNumbers)

	if actual != true {
		t.Errorf("%v should be valid", mockNumbers.GetSequence())
	}
}

func TestIsInvalid(t *testing.T) {
	mockNumbers := &mockDigitSeqence{
		sequence: []int{1, 2, 4},
	}

	actual := IsValid(mockNumbers)

	if actual != false {
		t.Errorf("%v should not be valid", mockNumbers.GetSequence())
	}
}
