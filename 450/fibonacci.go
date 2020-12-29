package main

import "fmt"

func main() {
	fmt.Println(GetNthFib(6))
}

func GetNthFib(n int) int {
	// Write your code here.
	if n == 2 {
		return 1
	} else if n == 1 {
		return 0
	}
	fmt.Println("first ", GetNthFib(n-1), "second ", GetNthFib(n-2))

	return GetNthFib(n-1) + GetNthFib(n-2)
}
