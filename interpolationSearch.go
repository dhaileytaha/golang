package main 

import "fmt" 

func interpolationSeach(array []int, key int) int {
		
	min, max := array[0], array[len(array)-1]

	low, high := 0, len(array)-1
	
	for {
		if key < min {
			return low
		}
		if key > 
	}	
}

func main() {
	
}