package container

import (
	"container/list"
)

// Stack implements a LIFO data structure, holding arbitrary data types
type Stack struct {
	*list.List
}

// NewStack creates an empty stack
func NewStack() *Stack {
	return &Stack{list.New()}
}

// Len returns the stack size
func (stack *Stack) Len() int {
	return stack.List.Len()
}

// Push inserts a new item in the stack
func (stack *Stack) Push(item interface{}) {
	stack.List.PushFront(item)
}

// Pop returns the front item in the stack, or nil if the stack is empty.
// Unlike Peek(), the item is removed from the stack
func (stack *Stack) Pop() interface{} {
	if e := stack.List.Front(); e != nil {
		return stack.List.Remove(e)
	}
	return nil
}

// Pop returns the front item in the stack, or nil if the stack is empty.
func (stack *Stack) Peek() interface{} {
	if e := stack.List.Front(); e != nil {
		return e.Value
	}
	return nil
}
