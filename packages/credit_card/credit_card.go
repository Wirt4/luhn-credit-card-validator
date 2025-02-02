package credit_card

import "main.go/packages/digit_container"

type CreditCard struct {
	container *digit_container.DigitContainer
}

func (c *CreditCard) SetSequence(sequence string) {
	c.container = digit_container.NewDigitContainer(sequence)
}

func (c CreditCard) GetSequence() []int {
	return c.container.GetSequence()
}

func (c *CreditCard) HasCorrectLength() bool {
	return c.container.IsSize(16)
}
