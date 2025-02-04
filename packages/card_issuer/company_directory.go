package card_issuer

import (
	"fmt"
	"strconv"
	"strings"
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

func ParseEntry(entry string) types.ProviderData {
	fmt.Println(entry)
	params := strings.SplitAfter(entry, " ")
	fmt.Println(params)
	name_builder := strings.Builder{}
	done := false
	var i int = 0
	for !done {
		fmt.Println(params[i])
		if string(params[i][0]) == "[" {
			fmt.Println("found bracket")
			done = true
			break
		}

		if isDigit(params[i][0]) {
			fmt.Printf("found number with '%v'", string(params[i][0]))
			done = true
			break
		}
		name_builder.WriteString(params[i])
		i++
	}
	name := strings.Trim(name_builder.String(), " ")
	fmt.Printf("testing for ending newline: %v%v\n", name, name)
	var iins []int
	if name == "VISA" {
		iins = []int{4}
	} else {
		iins = []int{55}
	}
	return types.ProviderData{
		Name:              name,
		IINs:              iins,
		MaxSequenceLength: 16,
		MinSequenceLength: 16,
	}

}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
