package digit_container

type DigitContainer struct {
	sequence []int
}

func NewDigitContainer(sequence string) *DigitContainer {
	container := &DigitContainer{
		sequence: []int{},
	}
	for _, v := range sequence {
		if v >= '0' && v <= '9' {
			container.sequence = append(container.sequence, int(v-'0'))
		}
	}
	return container
}

func (c *DigitContainer) GetSequence() []int {
	return c.sequence
}

func (c *DigitContainer) IsSize(length int) bool {
	return len(c.sequence) == length
}
