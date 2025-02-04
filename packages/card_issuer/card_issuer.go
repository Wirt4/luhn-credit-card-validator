package card_issuer

import "strconv"

type CardIssuer struct {
	SequenceLength int
	Issuer         string
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

func findIssuer(sequence []int, n *Node) *CardIssuer {
	if n == nil {
		return nil
	}

	if len(n.Children) == 0 {
		return n.Data
	}

	if len(sequence) == 0 {
		return nil
	}
	ndx, sequence := sequence[0], sequence[1:]

	return findIssuer(sequence, n.Children[ndx])
}

func NewCardIssuer(sequence []int, isInUS bool) *CardIssuer {
	n := buildTree(isInUS)
	issuer := findIssuer(sequence, n)
	if issuer == nil {
		return &CardIssuer{SequenceLength: -1, Issuer: "not found"}
	}
	return issuer
}

func buildTree(isInUS bool) *Node {
	//Another way to build the tree is to read from a file, which would make it more configurable
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
		if i == 55 && isInUS {
			provider += ": Diners Club U.S. and Canada"
		}
		tree.Insert(provider, i, 16)
	}
	for _, iin := range []int{65, 9792} {
		tree.Insert("Troy", iin, 16)
	}
	tree.Insert("UATP", 1, 15)
	tree.Insert("LankaPay", 357111, 16)
	if !isInUS {
		for _, iin := range []int{4026, 417500, 4508, 4844, 4913, 4917} {
			tree.Insert("Visa Electron", iin, 16)
		}
		for _, iin := range []int{5019, 4571} {
			tree.Insert("Dankort", iin, 16)
		}
	}
	tree.Insert("VISA", 4, 16)
	return tree.Root
}

type tree struct {
	Root *Node
}

func (t *tree) Insert(issuerName string, iin int, sequenceLength int) {
	t.Root = insertNode(strconv.Itoa(iin), &CardIssuer{SequenceLength: sequenceLength, Issuer: issuerName}, t.Root)
}

func newTree() *tree {
	return &tree{Root: &Node{Children: nil}}
}
