package container

import (
    "container/list"
)

type Queue struct {
    *list.List
}

func NewQueue() *Queue {
    return &Queue { list.New() }
}

func (q *Queue) Len() int {
    return q.List.Len()
}

func (q *Queue) Push(item interface{}) {
    q.List.PushBack(item)
}

func (q *Queue) Pop() interface{} {
    if e := q.List.Front(); e != nil {
        return q.List.Remove(e)
    }
    return nil
}

func (q *Queue) Peek() interface{} {
    if e := q.List.Front(); e != nil {
        return e.Value
    }
    return nil
}
