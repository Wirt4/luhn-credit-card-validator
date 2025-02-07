package card_issuer

import (
	"container/list"
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

func split(entry string) []string {
	temp := strings.Split(entry, " ")
	ans := []string{}
	for _, v := range temp {
		ans = append(ans, strings.Trim(v, " "))
	}
	return ans
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

type numberHandler struct {
	data []int
}

func (i *numberHandler) Range() []int {
	return i.data
}

func (i *numberHandler) stringisArray(s string) bool {
	return strings.Contains(s, "[") && strings.Contains(s, "]")
}

func (i *numberHandler) stringIsRange(s string) bool {
	return strings.Contains(s, "-")
}

func (i *numberHandler) parseNumber(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func (i *numberHandler) Set(s string) {
	q := newQueue()
	q.Add(s)
	numbers := []int{}
	for !q.isEmpty() {
		current := q.DeQueue()
		if i.stringisArray(current) {
			current = strings.TrimPrefix(current, "[")
			current = strings.TrimSuffix(current, "]")
			if !strings.Contains(current, ",") {
				q.Add(current)
				continue
			}
			contents := strings.Split(current, ", ")
			for _, v := range contents {
				q.Add(v)
			}
			continue
		}
		if i.stringIsRange(current) {
			r := strings.Split(current, "-")
			low := i.parseNumber(r[0])
			high := i.parseNumber(r[1])
			for j := low; j <= high; j++ {
				numbers = append(numbers, j)
			}
			continue
		}
		numbers = append(numbers, i.parseNumber(current))
	}
	i.data = numbers
}

func (i *numberHandler) GetLow() int {
	return 16
}

func (i *numberHandler) GetHigh() int {
	return 16
}

type queue struct {
	*list.List
}

func (q *queue) Add(v string) {
	q.PushBack(v)
}

func (q *queue) isEmpty() bool {
	return q.Len() == 0
}

func (q *queue) DeQueue() string {
	e := q.Front()
	q.List.Remove(e)
	return e.Value.(string)
}

// New is a new instance of a Queue
func newQueue() queue {
	return queue{list.New()}
}
