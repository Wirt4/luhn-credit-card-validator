package card_issuer

import (
	"strconv"

	"main.go/packages/types"
)

type tree struct {
	Root *types.Node
}

func newTree() *tree {
	return &tree{Root: &types.Node{Children: nil}}
}

func insertNode(iin string, issuer *types.CardIssuer, n *types.Node) *types.Node {
	if n == nil {
		n = &types.Node{Children: nil}
	}
	if iin == "" {
		n.Data = issuer
	} else {
		msd := int(iin[0] - '0')
		if len(iin) > 1 {
			iin = iin[1:]
		} else {
			iin = ""
		}
		if n.Children == nil {
			n.Children = make(map[int]*types.Node)
		}
		n.Children[msd] = insertNode(iin, issuer, n.Children[msd])
	}
	return n
}

func (t *tree) insert(issuerName string, iin int, sequenceLength int) {
	t.insertRange(issuerName, iin, sequenceLength, sequenceLength)
}

func (t *tree) insertRange(issuerName string, iin int, lowestSequenceLength int, highestSequenceLength int) {
	t.Root = insertNode(
		strconv.Itoa(iin),
		&types.CardIssuer{
			Min:    lowestSequenceLength,
			Max:    highestSequenceLength,
			Issuer: issuerName,
		},
		t.Root)
}
