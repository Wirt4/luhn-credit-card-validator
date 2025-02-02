package routing_number

import "main.go/packages/digit_container"

type RoutingNumber struct {
	container *digit_container.DigitContainer
}

func NewRoutingNumber(sequence string) *RoutingNumber {
	return &RoutingNumber{
		container: digit_container.NewDigitContainer(sequence),
	}
}

func (c *RoutingNumber) SetSequence(sequence string) {
	c.container = digit_container.NewDigitContainer(sequence)
}

func (c *RoutingNumber) GetSequence() []int {
	return c.container.GetSequence()
}

func (r *RoutingNumber) HasCorrectLength() bool {
	return r.container.IsSize(9)
}
