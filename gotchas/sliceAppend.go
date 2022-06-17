package main

import (
	"fmt"
)

func main() {
	var slice = make([]int, 3, 10)
	fmt.Println(slice) //[0,0,0]

	slice[0] = 1
	fmt.Println(slice) //[1 0 0]

	slice = append(slice, 4) // works, new element gets added to the slice
	fmt.Println(slice)       // [1 0 0 4]

	modifySlice(slice)
	fmt.Println(slice) //[11 0 3 4]
}

func modifySlice(s []int) {
	s[0] = 11

	s = append(s, 2) // new element gets added to the slice inside function, but not visible to caller function.
	s[2] = 3

	fmt.Println(s) //[11 0 3 4 2]
}
