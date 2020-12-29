package main

import "fmt"

func main() {
	arr := []int{2, 4, 1, 8, 5, 9, 0}
	fmt.Println(bubble(arr))
}

//time O(N*2) space O(1)
func bubble(array []int) []int {

	isSorted := false
	counter := 0
	for !isSorted {
		isSorted = true
		for i := 0; i < (len(array) - 1 - counter); i++ {

			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				isSorted = false
			}
		}
		counter++
	}
	return array
}

// for !isSorted {
// 		isSorted = true
// 		for i :=0; i<(len(array) - 1 - counter); i++ {
// 			if array[i] > array[i+1] {

// 					array[i], array[i+1] = array[i+1], array[i]
// 					isSorted = false

// 			}
// 		}
// 		counter++
// 	}
// 	return array
// }
