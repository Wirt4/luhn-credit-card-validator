package card_issuer

import (
	"strconv"
	"sync"

	"main.go/packages/types"
)

func GetInstance() *types.Node {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = buildTree()
		}
	}
	return singleInstance
}

var lock = &sync.Mutex{}
var singleInstance *types.Node

func buildTree() *types.Node {
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
