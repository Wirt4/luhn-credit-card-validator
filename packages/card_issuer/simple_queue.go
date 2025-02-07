package card_issuer

import "container/list"

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

func newQueue() queue {
	return queue{list.New()}
}

type queue struct {
	*list.List
}
