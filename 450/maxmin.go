//Time O(N)

package main

import "fmt"

// Pair is ...
type Pair struct {
	min int
	max int
}

func main() {
	arr := []int{22, 2, 50, 76, 91, 5, 58}

	fmt.Println(maxmin(arr, len(arr)))
}

func maxmin(arr []int, n int) Pair {
	var maxmin Pair

	if n == 1 {
		maxmin.min = arr[0]
		maxmin.max = arr[0]

		return maxmin
	}

	if arr[0] < arr[1] {
		maxmin.min = arr[0]
		maxmin.max = arr[1]
	} else {
		maxmin.min = arr[1]
		maxmin.max = arr[0]
	}

	for i := 2; i < n; i++ {
		if arr[i] < maxmin.min {
			maxmin.min = arr[i]
		} else if arr[i] > maxmin.max {
			maxmin.max = arr[i]
		}
	}
	return maxmin

}
