package issuer_visitor

import (
	"main.go/packages/types"
)

type Visitor struct {
	visited []types.CardIssuer
}

func (v *Visitor) Traverse(sequence []int, root *types.Node) {
	cur := root
	ndx := 0
	for cur != nil && ndx < len(sequence) && cur.Children != nil {
		cur = cur.Children[sequence[ndx]]
		ndx++
		if cur != nil && cur.Data != nil {
			v.visited = append(v.visited, *cur.Data)
		}
	}
}

func (v *Visitor) GetVisited() []types.CardIssuer {
	return v.visited
}
