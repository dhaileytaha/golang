package main

import "fmt"

func main() {
	arr := []int{3, 41, 11, 5, 9, 100}
	fmt.Println(sort(arr))

}

func sort(array []int) []int {
	currentIndex := 0

	for currentIndex < len(array)-1 {
		smallestIndex := currentIndex
		for i := currentIndex + 1; i < len(array); i++ {
			if array[smallestIndex] < array[i] {
				smallestIndex = i
			}
		}
		array[currentIndex], array[smallestIndex] = array[smallestIndex], array[currentIndex]
		currentIndex += 1
	}
	return array
}
