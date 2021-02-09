package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []string{"yo", "act", "flop", "tac", "foo", "cat", "oy", "olfp"}
	fmt.Println(anagramf(a))

}

func anagramf(words []string) [][]string {
	anagram := map[string][]string{}

	for _, word := range words {
		sorted := wordSorting(word)
		// fmt.Println("Word=", word)
		// fmt.Println("sorted=", sorted)

		anagram[sorted] = append(anagram[sorted], word)
		// fmt.Println("anagram=", anagram)
		// fmt.Println("anagram[sorted]=", anagram[sorted])

	}

	result := [][]string{}

	for _, group := range anagram {
		result = append(result, group)
	}
	return result
}

func wordSorting(word string) string {
	wordBytes := []byte(word)

	sort.Slice(wordBytes, func(i, j int) bool {

		return wordBytes[i] < wordBytes[j]

	})
	return string(wordBytes)
}
