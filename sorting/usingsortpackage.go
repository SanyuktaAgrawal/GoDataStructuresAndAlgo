package main

import (
	"fmt"
	"sort"
)

// Sort sort a string
func Sort(text string) string {
	runes := []rune(text)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func main() {

	name := "golangsortpackage"
	fmt.Println("--- Unsorted ---", name)
	name = Sort(name)
	fmt.Println("--- Sorted ---", name)
}
