package main

import "fmt"

func main() {
	source := "listen"
	target := "silent"

	fmt.Println(isAnagram(source, target))
}

func isAnagram(source, target string) bool {
	if len(source) != len(target) {
		return false
	}

	anagramMap := make(map[string]int)

	for _, val := range source {
		anagramMap[string(val)]++
	}

	for _, val := range target {
		anagramMap[string(val)]--
	}

	for val, i := range anagramMap {
		if i != 0 {
			return false
		}
	}

	return true
}
