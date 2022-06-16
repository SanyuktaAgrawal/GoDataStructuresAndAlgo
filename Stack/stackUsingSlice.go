package main

import (
	"fmt"
)

func Len(stack []string) int {
	return len(stack)
}

func Push(stack []string, value string) {
	fmt.Println(cap(stack))
	fmt.Println(stack)
	stack = append(stack, value)
	fmt.Println(stack)
}

func Pop(stack []string) string {
	fmt.Println(Len(stack))
	fmt.Println(stack)
	n := len(stack) - 1

	return stack[n]
}

func main() {
	//var stack []string
	var stack = make([]string, 5)
	Push(stack, "world!")
	Push(stack, "Hello ")

	fmt.Print(Pop(stack))
}
