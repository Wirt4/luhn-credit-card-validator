package number_parser

import (
	"strconv"
	"strings"
	"unicode"

	"main.go/packages/simple_queue"
	"main.go/packages/types"
)

type NumberParser struct {
	data []int
}

func (i *NumberParser) Range() []int {
	return i.data
}

func (i *NumberParser) stringisArray(s string) bool {
	return strings.Contains(s, "[") && strings.Contains(s, "]")
}

func (i *NumberParser) TrimBrackets(s string) string {
	s = strings.TrimPrefix(s, "[")
	return strings.TrimSuffix(s, "]")
}

func (i *NumberParser) stringIsRange(s string) bool {
	return strings.Contains(s, "-")
}

func (i *NumberParser) parseNumber(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func (i *NumberParser) parseSlice(s string) []string {
	s = i.TrimBrackets(s)
	if !strings.Contains(s, ",") {
		return []string{s}
	}
	return strings.Split(s, ", ")
}

func (i *NumberParser) parseRange(s string) (int, int) {
	r := strings.Split(s, "-")
	low := i.parseNumber(r[0])
	high := i.parseNumber(r[1])
	return low, high
}

func (i *NumberParser) Set(s string) {
	q := simple_queue.NewQueue()
	q.Add(s)
	i.data = []int{}
	for !q.IsEmpty() {
		current := q.DeQueue()
		if i.stringisArray(current) {
			contents := i.parseSlice(current)
			for _, v := range contents {
				q.Add(v)
			}
		} else if i.stringIsRange(current) {
			low, high := i.parseRange(current)
			for j := low; j <= high; j++ {
				i.data = append(i.data, j)
			}
		} else {
			i.data = append(i.data, i.parseNumber(current))
		}
	}
}

func (i *NumberParser) GetLow() int {
	return i.data[0]
}

func (i *NumberParser) GetHigh() int {
	return i.data[len(i.data)-1]
}

func ParseEntry(entry string) types.ProviderData {
	params := split(entry)
	provider_name := strings.Builder{}
	iins := &NumberParser{}
	sequence_range := &NumberParser{}
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

func split(entry string) []string {
	var i = 1
	var j = 0
	splitIndeces := [2]int{}
	for i < len(entry)-1 {
		seg := entry[i-1 : i+2]
		if isJuncture(seg) {
			splitIndeces[j] = i
			j++
		}
		i++
	}

	name := entry[:splitIndeces[0]]
	iins := entry[splitIndeces[0]+1 : splitIndeces[1]]
	prefixes := entry[splitIndeces[1]+1:]

	return []string{name, iins, prefixes}
}

func isJuncture(segment string) bool {
	if len(segment) < 3 {
		return false
	}
	if segment[0] == ',' {
		return false
	}
	if segment[1] != ' ' {
		return false
	}
	if segment[0] == ']' || isDigit(segment[0]) {
		return true
	}
	return segment[2] == '[' || isDigit(segment[2])
}

func isDigit(ch byte) bool {
	return unicode.IsDigit(rune(ch))
}
