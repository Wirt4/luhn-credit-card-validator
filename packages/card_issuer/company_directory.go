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

func GetInstance() (*types.Node, error) {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			instance, err := buildTree()
			if err != nil {
				return nil, err
			}
			singleInstance = instance
		}
	}
	return singleInstance, nil
}

var lock = &sync.Mutex{}
var singleInstance *types.Node

func buildTree() (*types.Node, error) {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/providers.txt")
	if err != nil {
		fmt.Println("Error opening file")
		return nil, err
	}
	defer file.Close()
	tree := newTree()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entry := ParseEntry(scanner.Text())
		for _, iin := range entry.IINs {
			min := entry.MinSequenceLength
			max := entry.MaxSequenceLength
			tree.insertRange(entry.Name, iin, min, max)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return tree.Root, nil
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
	params := split(entry)
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
