package credit_card

import (
	"main.go/packages/card_issuer"
	"main.go/packages/issuer_visitor"
	"main.go/packages/types"
)

type CreditCard struct {
	sequence []int
	issuers  []types.CardIssuer
}

func NewCreditCard() *CreditCard {
	return &CreditCard{}
}

func (card *CreditCard) SetSequence(sequence string) error {
	for _, v := range sequence {
		if v >= '0' && v <= '9' {
			card.sequence = append(card.sequence, int(v-'0'))
		}
	}
	if len(card.issuers) == 0 {
		issuers, error := card_issuer.GetCardIssuers(card.sequence, &issuer_visitor.Visitor{})
		if error != nil {
			return error
		}
		card.issuers = issuers
	}
	return nil
}

func (card CreditCard) GetSequence() []int {
	return card.sequence
}

func (card *CreditCard) HasCorrectLength() bool {
	filtered_issuers := []types.CardIssuer{}
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
