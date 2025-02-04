package card_issuer

import (
	"strconv"
)

type CardIssuer struct {
	Min    int
	Max    int
	Issuer string
}

func GetCardIssuers(sequence []int) []CardIssuer {
	issuer_directory := buildTree()
	v := visitor{sequence: sequence, ndx: 0, cur: issuer_directory, visited: []CardIssuer{}}
	v.traverse()
	return v.visited
}

type node struct {
	Children map[int]*node
	Data     *CardIssuer
}

func buildTree() *node {
	tree := newTree()
	tree.insert("VISA", 4, 16)
	for _, iin := range []int{34, 37} {
		tree.insert("American Express", iin, 15)
	}
	tree.insert("China T-Union", 31, 19)
	for _, iin := range []int{60, 65, 81, 82, 508, 353, 356} {
		tree.insert("RuPay", iin, 16)
	}
	tree.insert("BORICA-BANKCARD", 25, 16)
	for i := 2221; i <= 2720; i++ {
		tree.insert("Mastercard", i, 16)
	}
	for i := 51; i <= 59; i++ {
		var provider = "Mastercard"
		if i == 55 {
			provider += ": Diners Club U.S. and Canada"
		}
		tree.insert(provider, i, 16)
	}
	for _, iin := range []int{65, 9792} {
		tree.insert("Troy", iin, 16)
	}
	tree.insert("UATP", 1, 15)
	tree.insert("LankaPay", 357111, 16)

	for _, iin := range []int{4026, 417500, 4508, 4844, 4913, 4917} {
		tree.insert("Visa Electron", iin, 16)
	}
	for _, iin := range []int{5019, 4571} {
		tree.insert("Dankort", iin, 16)
	}

	tree.insertRange("China Union Pay", 62, 16, 19)
	return tree.Root
}

type visitor struct {
	sequence []int
	ndx      int
	cur      *node
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

type tree struct {
	Root *node
}

func newTree() *tree {
	return &tree{Root: &node{Children: nil}}
}

func (t *tree) insert(issuerName string, iin int, sequenceLength int) {
	t.insertRange(issuerName, iin, sequenceLength, sequenceLength)
}

func (t *tree) insertRange(issuerName string, iin int, start int, end int) {
	t.Root = insertNode(
		strconv.Itoa(iin),
		&CardIssuer{
			Min:    start,
			Max:    end,
			Issuer: issuerName,
		},
		t.Root)
}

func insertNode(iin string, issuer *CardIssuer, n *node) *node {
	if n == nil {
		n = &node{Children: nil}
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
			n.Children = make(map[int]*node)
		}
		n.Children[msd] = insertNode(iin, issuer, n.Children[msd])
	}
	return n
}
