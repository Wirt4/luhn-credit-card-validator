package luhn

import (
	"main.go/packages/interfaces"
)

func IsValid(sequence interfaces.DigitSequence) bool {
	if !sequence.HasCorrectLength() {
		return false
	}
	var s = sequence.GetSequence()
	if len(s) < 2 {
		return false
	}

	lastIndex := len(s) - 1
	checkDigit := s[lastIndex]
	s = s[:lastIndex]
	var sum = 0

	for i, v := range s {
		if i%2 == 0 {
			sum += v
		} else {
			sum += sumDigit(v)
		}
	}

	return modulateDigit(sum) == checkDigit
}

func sumDigit(digit int) int {
	digit *= 2

	if digit > 9 {
		return addDigits(digit)
	}

	return digit
}

func addDigits(digit int) int {
	sum := 0

	for digit > 0 {
		sum += digit % 10
		digit /= 10
	}

	return sum
}

func modulateDigit(digit int) int {
	return (10 - digit%10) % 10
}
