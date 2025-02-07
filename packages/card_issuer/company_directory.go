package card_issuer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"unicode"

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
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/providers.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	tree := newTree()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entry := ParseEntry(scanner.Text())
		for _, iin := range entry.IINs {
			fmt.Printf("name %v\nIINs %v\nmin length %v\nmax length %v\n\n", entry.Name, entry.IINs, entry.MinSequenceLength, entry.MaxSequenceLength)
			tree.insertRange(entry.Name, iin, entry.MinSequenceLength, entry.MaxSequenceLength)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return tree.Root
}

func temp() *types.Node {
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

func split(entry string) []string {
	var i = 1
	var j = 0
	splitIndeces := [2]int{}
	for i < len(entry)-1 {
		if entry[i] == ' ' && entry[i-1] != ',' && ((entry[i-1] == ']' || unicode.IsDigit(rune(entry[i-1]))) || (entry[i+1] == '[' || unicode.IsDigit(rune(entry[i+1])))) {
			splitIndeces[j] = i
			j++
		}
		i++
	}
	return []string{entry[:splitIndeces[0]], entry[splitIndeces[0]+1 : splitIndeces[1]], entry[splitIndeces[1]+1:]}
}

func ParseEntry(entry string) types.ProviderData {
	params := split(entry) //splits on all spaces
	provider_name := strings.Builder{}
	iins := &numberHandler{}
	sequence_range := &numberHandler{}
	sequence_range.Set(params[len(params)-1])
	iins.Set(params[len(params)-2])

	for i := 0; i < len(params)-2; i++ {
		if i != 0 {
			provider_name.WriteString(" ")
		}
		provider_name.WriteString(params[i])

	}

	return types.ProviderData{
		Name:              provider_name.String(),
		IINs:              iins.Range(),
		MaxSequenceLength: sequence_range.GetHigh(),
		MinSequenceLength: sequence_range.GetLow(),
	}
}
