package container

import (
	"container/list"
)

// Queue implements a FIFO data structure, holding arbitrary data types
type Queue struct {
	*list.List
}

// NewQueue creates an empty queue
func NewQueue() *Queue {
	return &Queue{list.New()}
}

// Len returns the queue size
func (q *Queue) Len() int {
	return q.List.Len()
}

// Push inserts a new item in the queue
func (q *Queue) Push(item interface{}) {
	q.List.PushBack(item)
}

// Pop returns the front item in the queue, or nil if the queue is empty.
// Unlike Peek(), the item is removed from the queue
func (q *Queue) Pop() interface{} {
	if e := q.List.Front(); e != nil {
		return q.List.Remove(e)
	}
	return nil
}

// Pop returns the front item in the queue, or nil if the queue is empty.
func (q *Queue) Peek() interface{} {
	if e := q.List.Front(); e != nil {
		return e.Value
	}
	return nil
}
