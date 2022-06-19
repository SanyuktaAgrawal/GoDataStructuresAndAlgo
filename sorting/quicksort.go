package main

import (
	"fmt"
	"math/rand"
)

func main() {

	slice := []int{243, -412, 380, -312, 925, 158, -46, 177, 22, -482, 273, 217, 514, -392, 424}
	fmt.Println("--- Unsorted ---", slice)
	quicksort(slice)
	fmt.Println("--- Sorted ---", slice)
}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)
	fmt.Println("--- pivot ---", pivot)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	// place pivot at correct position by swapping a[left] or a[right] (or pivot)
	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}
