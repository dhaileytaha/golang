package main

import "fmt"

// O(n) time and o(1) space

func main() {
	str := "abcdcba"
	fmt.Println(palin(str))
}

func palin(str string) bool {

	for i := 0; i < len(str); i++ {
		j := len(str) - 1 - i

		if str[i] != str[j] {
			return false
		}
	}
	return true
}
