package credit_card

import "main.go/packages/card_issuer"

type CreditCard struct {
	sequence []int
	issuers  []card_issuer.CardIssuer
}

func NewCreditCard() *CreditCard {
	return &CreditCard{}
}

func (card *CreditCard) SetSequence(sequence string) {
	for _, v := range sequence {
		if v >= '0' && v <= '9' {
			card.sequence = append(card.sequence, int(v-'0'))
		}
	}
	card.issuers = card_issuer.GetCardIssuers(card.sequence)
}

func (card CreditCard) GetSequence() []int {
	return card.sequence
}

func (card *CreditCard) HasCorrectLength() bool {
	filtered_issuers := []card_issuer.CardIssuer{}
	var isCorrect bool = false
	for _, issuer := range card.issuers {
		if issuer.Min <= len(card.sequence) && len(card.sequence) <= issuer.Max {
			filtered_issuers = append(filtered_issuers, issuer)
			isCorrect = true
		}
	}
	card.issuers = filtered_issuers
	return isCorrect
}
