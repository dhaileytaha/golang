package main

import "fmt"

func main() {
	arr := []int{5, 2, 6, 3, 1, 4}
	fmt.Println(arr)

	fmt.Println(rev(arr, 0, len(arr)-1))

}

// [1,2,4,5,6,78,8]

func rev(arr []int, start int, end int) []int {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
	return arr
}
