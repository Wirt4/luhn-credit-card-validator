package credit_card

import "main.go/packages/card_issuer"

type CreditCard struct {
	sequence []int
	issuer   *card_issuer.CardIssuer
}

func NewCreditCard() *CreditCard {
	return &CreditCard{issuer: &card_issuer.CardIssuer{
		SequenceLength: 16,
	}}
}

func (card *CreditCard) SetSequence(sequence string) {
	for _, v := range sequence {
		if v >= '0' && v <= '9' {
			card.sequence = append(card.sequence, int(v-'0'))
		}
	}
	//create the provider
}

func (card CreditCard) GetSequence() []int {
	return card.sequence
}

func (card *CreditCard) HasCorrectLength() bool {
	return len(card.sequence) == card.issuer.SequenceLength
}
