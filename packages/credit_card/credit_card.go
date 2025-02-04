package credit_card

import "main.go/packages/card_issuer"

type CreditCard struct {
	sequence []int
	issuers  []*card_issuer.CardIssuer
}

func NewCreditCard() *CreditCard {
	return &CreditCard{issuers: []*card_issuer.CardIssuer{{
		Min: 16,
		Max: 16,
	}}}
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
	return len(card.sequence) >= card.issuers[0].Min && len(card.sequence) <= card.issuers[0].Max
}
