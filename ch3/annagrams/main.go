package main

import "fmt"

func main() {
	fmt.Printf("Is anagram: %v\n", isAnagrams("makapaka", "makapaka"))
}

func isAnagrams(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	lettersCOne := countLetters(s1)
	lettersCTwo := countLetters(s2)

	for r, n := range lettersCOne {
		if n != lettersCTwo[r] {
			return false
		}
	}
	return true
}

func countLetters(s string) map[rune]int {
	counter := make(map[rune]int)
	for _, v := range s {
		counter[v]++
	}
	return counter
}
