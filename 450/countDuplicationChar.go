package main

import (
	"fmt"
)

func main() {
	str := "test"

	fmt.Println(count1(str))
}

func count1(str string) map[string]int {
	countLetter := make(map[string]int)

	for i := 0; i < len(str); i++ {
		countLetter[string(str[i])] = 1

		for index := range countLetter {
			if index == string(str[i]) {
				fmt.Println(index)
				fmt.Println(string(str[i]))
				countLetter[string(str[i])]++
			}
		}
		// fmt.Println(sc)
	}
	return countLetter
}
