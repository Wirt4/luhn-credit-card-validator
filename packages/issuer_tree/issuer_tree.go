package issuer_tree

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"

	"main.go/packages/number_parser"
	"main.go/packages/types"
)

type Tree struct {
	Root *types.Node
}

func NewTree() *Tree {
	return &Tree{Root: &types.Node{Children: nil}}
}

func GetInstance() (*types.Node, error) {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		instance, err := scanFileIntoTree()
		if err != nil {
			return nil, err
		}
		singleInstance = instance

	}
	return singleInstance, nil
}

var lock = &sync.Mutex{}
var singleInstance *types.Node

func (t *Tree) InsertRange(issuerName string, iin int, lowestSequenceLength int, highestSequenceLength int) {
	t.Root = insertNode(
		strconv.Itoa(iin),
		&types.CardIssuer{
			Min:    lowestSequenceLength,
			Max:    highestSequenceLength,
			Issuer: issuerName,
		},
		t.Root)
}

func scanFileIntoTree() (*types.Node, error) {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/providers.txt")
	if err != nil {
		fmt.Println("Error opening file")
		return nil, err
	}
	defer file.Close()

	tree := NewTree()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entry := number_parser.ParseEntry(scanner.Text())
		for _, iin := range entry.IINs {
			min := entry.MinSequenceLength
			max := entry.MaxSequenceLength
			tree.InsertRange(entry.Name, iin, min, max)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return tree.Root, nil
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
