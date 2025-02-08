package simple_queue

import "container/list"

func (q *queue) Add(v string) {
	q.PushBack(v)
}

func (q *queue) IsEmpty() bool {
	return q.Len() == 0
}

func (q *queue) DeQueue() string {
	e := q.Front()
	q.List.Remove(e)
	return e.Value.(string)
}

func NewQueue() queue {
	return queue{list.New()}
}

type queue struct {
	*list.List
}
