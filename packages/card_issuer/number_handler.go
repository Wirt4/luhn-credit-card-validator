package card_issuer

import (
	"strconv"
	"strings"
)

type numberHandler struct {
	data []int
}

func (i *numberHandler) Range() []int {
	return i.data
}

func (i *numberHandler) stringisArray(s string) bool {
	return strings.Contains(s, "[") && strings.Contains(s, "]")
}

func (i *numberHandler) TrimBrackets(s string) string {
	s = strings.TrimPrefix(s, "[")
	return strings.TrimSuffix(s, "]")
}

func (i *numberHandler) stringIsRange(s string) bool {
	return strings.Contains(s, "-")
}

func (i *numberHandler) parseNumber(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func (i *numberHandler) parseSlice(s string) []string {
	s = i.TrimBrackets(s)
	if !strings.Contains(s, ",") {
		return []string{s}
	}
	return strings.Split(s, ", ")
}

func (i *numberHandler) parseRange(s string) (int, int) {
	r := strings.Split(s, "-")
	low := i.parseNumber(r[0])
	high := i.parseNumber(r[1])
	return low, high
}

func (i *numberHandler) Set(s string) {
	q := newQueue()
	q.Add(s)
	i.data = []int{}
	for !q.isEmpty() {
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

func (i *numberHandler) GetLow() int {
	return 16
}

func (i *numberHandler) GetHigh() int {
	return 16
}
