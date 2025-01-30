package luhn

import (
	"main.go/packages/interfaces"
)

func IsValid(sequence interfaces.DigitSequence) bool {
	stack := newStackFromSequence(sequence.GetSequence())
	check_digit := stack.pop()
	return check_digit == calculateCheckDigit(stack)
}

func calculateCheckDigit(stack *stack) int {
	odd := stack.size()%2 != 0
	sum := 0
	for !stack.isEmpty() {
		digit := stack.pop()
		if odd {
			sum += digit
		} else {
			sum += sumDigit(digit)
		}
		odd = !odd
	}
	return (10 - (sum % 10)) % 10
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

type stack struct {
	sequence []int
}

func (s *stack) size() int {
	return len(s.sequence)
}

func newStackFromSequence(sequence []int) *stack {
	s := &stack{
		sequence: sequence,
	}
	return s
}

func (s *stack) isEmpty() bool {
	return len(s.sequence) == 0
}

func (s *stack) pop() int {
	if len(s.sequence) == 0 {
		return -1
	}
	lastIndex := len(s.sequence) - 1
	lastDigit := s.sequence[lastIndex]
	s.sequence = s.sequence[:lastIndex]
	return lastDigit
}
