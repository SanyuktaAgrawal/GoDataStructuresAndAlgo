package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []string{"input", "age", "college", "school", "address"}
	fmt.Println("--- Unsorted ---", input)
	sort.Strings(input)
	fmt.Println("--- Sorted ---", input)
}
