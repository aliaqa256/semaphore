package queue

import (
	"fmt"
	"sync"
	
)

type Queue struct {
	sync.Mutex
	items []int
}

func (q *Queue) Enqueue(i int) {
	q.Lock()
	defer q.Unlock()
	q.items = append(q.items, i)
}

func (q *Queue) Dequeue() int {
	q.Lock()
	defer q.Unlock()
	if q.isEmpty() {
		return -1
	}
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func (q *Queue) IsEmpty() bool {
	q.Lock()
	defer q.Unlock()
	return q.isEmpty()
}

func (q *Queue) isEmpty() bool {

	return len(q.items) == 0
}

func (q *Queue) Peek() int {
	if q.IsEmpty() {
		return -1
	}
	q.Lock()
	defer q.Unlock()
	return q.items[0]
}

func (q *Queue) Size() int {
	q.Lock()
	defer q.Unlock()
	return len(q.items)
}

func (q *Queue) Print() {

	q.Lock()
	defer q.Unlock()

	for _, v := range q.items {
		fmt.Print(v, " ")
	}
	fmt.Println()
}