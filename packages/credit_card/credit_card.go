package credit_card

// TODO: Implement the CreditCard struct and functions based off the interface "NumberSequence"
type CreditCard struct{}

func NewCreditCard() CreditCard {
	return CreditCard{}
}

func (c *CreditCard) SetSequence(sequence string) {

}

func (c *CreditCard) GetSequence() []int {
	return []int{4, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
}
