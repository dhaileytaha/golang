package main

import "fmt"

func main() {
	a := []int{1, 2, 2, 3, 45, 2, 1}

	final := common(a)
	fmt.Println(final)

}

func common(a []int) int {
	count := 0
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				count++
			}
		}
	}
	return count
}
