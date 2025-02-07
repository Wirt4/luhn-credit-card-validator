package luhn

import (
	"main.go/packages/interfaces"
)

type LuhnValidator struct{}

func (v *LuhnValidator) IsValid(sequence interfaces.DigitSequence) (bool, error) {
	if !sequence.HasCorrectLength() {
		return false, nil
	}
	var s = sequence.GetSequence()
	if len(s) < 2 {
		return false, nil
	}

	lastIndex := len(s) - 1
	checkDigit := s[lastIndex]
	s = s[:lastIndex]
	var sum = 0

	for i, j := range s {
		if i%2 == 0 {
			sum += j
		} else {
			sum += v.sumDigit(j)
		}
	}

	return v.modulateDigit(sum) == checkDigit, nil
}

func (v *LuhnValidator) sumDigit(digit int) int {
	digit *= 2

	if digit > 9 {
		return v.addDigits(digit)
	}

	return digit
}

func (v *LuhnValidator) addDigits(digit int) int {
	sum := 0

	for digit > 0 {
		sum += digit % 10
		digit /= 10
	}

	return sum
}

func (v *LuhnValidator) modulateDigit(digit int) int {
	return (10 - digit%10) % 10
}
