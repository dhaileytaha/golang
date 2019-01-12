package main 

import "fmt"

func linearSearch(datalist []int, value int) bool {
	for _, item := range datalist {
		if item == value {
			return true
		}
	}
	return false
}

func main() {
	item := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(linearSearch(item, 3))
}