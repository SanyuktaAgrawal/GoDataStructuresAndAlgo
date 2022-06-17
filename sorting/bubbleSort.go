package main

import "fmt"

//A function to implement bubble sort
func bubbleSort(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	bubbleSort(arr, len(arr))
	fmt.Println("Sorted array: ", arr)
}
