package container

import (
    "container/list"
)

type Stack struct {
    *list.List
}

func NewStack() *Stack {
    return &Stack { list.New() }
}

func (stack *Stack) Len() int {
    return stack.List.Len()
}

func (stack *Stack) Push(item interface{}) {
    stack.List.PushFront(item)
}

func (stack *Stack) Pop() interface{} {
    if e := stack.List.Front(); e != nil {
        return stack.List.Remove(e)
    }
    return nil
}

func (stack *Stack) Peek() interface{} {
    if e := stack.List.Front(); e != nil {
        return e.Value
    }
    return nil
}
