package card_issuer

import (
	"strconv"
)

type CardIssuer struct {
	Min    int
	Max    int
	Issuer string
}

type Node struct {
	Children map[int]*Node
	Data     *CardIssuer
}

func insertNode(iin string, issuer *CardIssuer, n *Node) *Node {
	if n == nil {
		n = &Node{Children: nil}
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
			n.Children = make(map[int]*Node)
		}
		n.Children[msd] = insertNode(iin, issuer, n.Children[msd])
	}
	return n
}

func GetCardIssuers(sequence []int) []CardIssuer {
	n := buildTree()
	v := visitor{sequence: sequence, ndx: 0, cur: n, visited: []CardIssuer{}}
	v.traverse()
	return v.visited
}

type visitor struct {
	sequence []int
	ndx      int
	cur      *Node
	visited  []CardIssuer
}

func (v *visitor) traverse() {
	for v.cur != nil && v.ndx < len(v.sequence) && v.cur.Children != nil {
		v.cur = v.cur.Children[v.sequence[v.ndx]]
		v.ndx++
		if v.cur != nil && v.cur.Data != nil {
			v.visited = append(v.visited, *v.cur.Data)
		}
	}
}

func buildTree() *Node {
	tree := newTree()
	tree.Insert("VISA", 4, 16)
	for _, iin := range []int{34, 37} {
		tree.Insert("American Express", iin, 15)
	}
	tree.Insert("China T-Union", 31, 19)
	for _, iin := range []int{60, 65, 81, 82, 508, 353, 356} {
		tree.Insert("RuPay", iin, 16)
	}
	tree.Insert("BORICA-BANKCARD", 25, 16)
	for i := 2221; i <= 2720; i++ {
		tree.Insert("Mastercard", i, 16)
	}
	for i := 51; i <= 59; i++ {
		var provider = "Mastercard"
		if i == 55 {
			provider += ": Diners Club U.S. and Canada"
		}
		tree.Insert(provider, i, 16)
	}
	for _, iin := range []int{65, 9792} {
		tree.Insert("Troy", iin, 16)
	}
	tree.Insert("UATP", 1, 15)
	tree.Insert("LankaPay", 357111, 16)

	for _, iin := range []int{4026, 417500, 4508, 4844, 4913, 4917} {
		tree.Insert("Visa Electron", iin, 16)
	}
	for _, iin := range []int{5019, 4571} {
		tree.Insert("Dankort", iin, 16)
	}

	tree.Insert("VISA", 4, 16)

	tree.InsertRange("China Union Pay", 62, 16, 19)
	return tree.Root
}

type tree struct {
	Root *Node
}

func (t *tree) Insert(issuerName string, iin int, sequenceLength int) {
	t.InsertRange(issuerName, iin, sequenceLength, sequenceLength)
}

func (t *tree) InsertRange(issuerName string, iin int, start int, end int) {
	t.Root = insertNode(
		strconv.Itoa(iin),
		&CardIssuer{
			Min:    start,
			Max:    end,
			Issuer: issuerName,
		},
		t.Root)
}

func newTree() *tree {
	return &tree{Root: &Node{Children: nil}}
}
