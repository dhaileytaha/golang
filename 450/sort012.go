package main

import "fmt"

func main() {
	arr := []int{0, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 1}
	fmt.Println(sort(arr))
}

func sort(arr []int) []int {

	c0 := 0
	c1 := 0
	c2 := 0

	i := 0

	for i = 0; i < len(arr); i++ {

		switch arr[i] {
		case 0:
			c0++
			break
		case 1:
			c1++
			break
		case 2:
			c2++
			break
		}
	}

	// Update the array
	i = 0

	// Store all the 0s in the beginning
	for c0 > 0 {
		arr[i] = 0
		i++
		c0--

	}

	// Then all the 1s
	for c1 > 0 {
		arr[i] = 1
		i++
		c1--
	}

	// Finally all the 2s
	for c2 > 0 {
		arr[i] = 2
		i++
		c2--
	}

	// Print the sorted array
	return arr

}
