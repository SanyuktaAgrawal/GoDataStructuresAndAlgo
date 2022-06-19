package main

import (
	"container/list"
	"fmt"
)

type stack struct {
	stack *list.List
}

func (c *stack) Push(value string) {
	c.stack.PushFront(value)
}

func (c *stack) Pop() {
	if c.stack.Len() > 0 {
		ele := c.stack.Front()
		c.stack.Remove(ele)
	}
}

func main() {
	stack := &stack{
		stack: list.New(),
	}

	stack.Push("A")
	stack.Push("B")

	stack.Pop()

	fmt.Println(stack)
}
