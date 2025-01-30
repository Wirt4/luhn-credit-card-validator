package credit_card

// TODO: Implement the CreditCard struct and functions based off the interface "NumberSequence"
type CreditCard struct {
	sequence []int
	valid    bool
}

func NewCreditCard() CreditCard {
	return CreditCard{
		valid: true,
	}
}

func (c *CreditCard) SetSequence(sequence string) {
	for _, v := range sequence {
		if isNumber(v) {
			c.sequence = append(c.sequence, int(v-'0'))
		}
	}

}

func (c *CreditCard) GetSequence() []int {
	return c.sequence
}

func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

func (c *CreditCard) HasCorrectLength() bool {
	return len(c.sequence) == 16
}
