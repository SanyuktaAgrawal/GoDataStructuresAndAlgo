package main

import (
	"fmt"
)

func main() {
	var n *int
	var i interface{}
	fmt.Println(n == nil) //true
	fmt.Println(i == nil) //true
	fmt.Println(n == i)   // false (type different)
}
