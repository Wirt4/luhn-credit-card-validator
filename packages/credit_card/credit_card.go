package credit_card

type CreditCard struct {
	sequence []int
}

func (c *CreditCard) SetSequence(sequence string) {
	for _, v := range sequence {
		if v >= '0' && v <= '9' {
			c.sequence = append(c.sequence, int(v-'0'))
		}
	}
}

func (c CreditCard) GetSequence() []int {
	return c.sequence
}

func (c *CreditCard) HasCorrectLength() bool {
	return len(c.sequence) == 16
}
