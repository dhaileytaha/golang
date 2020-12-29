package main

import "fmt"

func main() {
	arr := []int{2, 4, 1, 8, 5, 9, 0}
	fmt.Println(geti(arr))
}

//time O(N*2) space O(1)
func geti(array []int) []int {
	for i := range array {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			array[j], array[j-1] = array[j-1], array[j]
		}
	}
	return array
}
