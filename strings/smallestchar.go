package main

import "fmt"

func main() {

	inputStr := "iaminputstring"
	min := findSmallestChar(inputStr)
	fmt.Println("Smallest char is: ", string(min))

}

func findSmallestChar(input string) rune {
	inputRunes := []rune(input)
	min := inputRunes[0]
	for _, value := range inputRunes {
		if value < min {
			min = value
		}
	}
	return min
}
