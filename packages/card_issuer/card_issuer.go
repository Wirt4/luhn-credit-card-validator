package card_issuer

import (
	"main.go/packages/interfaces"
	"main.go/packages/types"
)

func GetCardIssuers(sequence []int, visitor interfaces.Visitor) []types.CardIssuer {
	issuer_directory := GetInstance()
	visitor.Traverse(sequence, issuer_directory)
	return visitor.GetVisited()
}
