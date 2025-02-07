package card_issuer

import (
	"main.go/packages/interfaces"
	"main.go/packages/types"
)

func GetCardIssuers(sequence []int, visitor interfaces.Visitor) ([]types.CardIssuer, error) {
	singleInstance, error := GetInstance()
	if error != nil {
		return nil, error
	}
	visitor.Traverse(sequence, singleInstance)
	return visitor.GetVisited(), nil
}
