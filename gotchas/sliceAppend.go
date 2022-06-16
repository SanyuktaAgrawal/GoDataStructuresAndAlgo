package main

import (
	"fmt"
)

func main() {
	var slice = make([]int, 3, 10) //len=3,cap=5
	fmt.Println(slice)             //[0,0,0]

	slice[0] = 1
	fmt.Println(slice) //[1 0 0]

	slice = append(slice, 4)
	fmt.Println(slice) // [1 0 0 4]

	AddOne(slice)
	fmt.Println(slice) //[11 0 3 4]
}

func AddOne(s []int) {
	s[0] = 11

	//when the new elements append to the slice a new array is created.
	//Now, the elements present in the existing array are copied into a new array and return a new slice reference to this array (new array).
	s = append(s, 2)
	s[2] = 3

	fmt.Println(s) //[11 0 3 4 2]
}
