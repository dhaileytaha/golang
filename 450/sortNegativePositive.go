package main

import "fmt"

func main() {
	arr := []int{-12, 11, -13, -5, 6, -7, 5, -3, -6}
	fmt.Println("input array:", arr)
	fmt.Println("output array", sort(arr))
}

func sort(arr []int) []int {
	i := 0
	j := len(arr) - 1
	for i = range arr {

		if arr[i] > 0 && arr[j] < 0 && i < j {

			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}

		if arr[i] < 0 && i < j {
			i++
		}
		if arr[j] > 0 && i < j {
			j--
		}
	}
	return arr
}
