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
	// Arrange
	mockNumbers := &mockDigitSeqence{
		sequence: []int{1, 2, 5},
	}

	actual := IsValid(mockNumbers)

	if actual != true {
		t.Errorf("%v should be valid", mockNumbers.GetSequence())
	}

}
