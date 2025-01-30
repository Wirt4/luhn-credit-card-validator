package luhn

import "main.go/packages/interfaces"

func IsValid(sequence interfaces.DigitSequence) bool {
	seq := sequence.GetSequence()
	lastDigit := seq[len(seq)-1]
	return lastDigit == 5
}
