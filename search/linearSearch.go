package main

import "fmt"

func search(arr []int, element int) (int, bool) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == element {
			return i, true
		}
	}
	return 0, false
}

func main() {
	arr := []int{2, 3, 4, 10, 40}
	element := 40

	index, found := search(arr, element)
	if !found {
		fmt.Printf("Element is not present in array")
	} else {
		fmt.Printf("Element is present at index %d", index)
	}
}
