package main 

import "fmt"

func binarySearch(datalist []int, key int) bool {

    low := 0
    high := len(datalist) - 1

    for low <= high {
        median := (low + high) / 2

        if datalist[median] < key {
            low = median + 1
        } else {
            high = median - 1
        }
    }
        if low == len(datalist) || datalist[low] != key {
            return false
        }
        return true
    
}

func main() {
      list := []int{1,2,3,4,5,6,7,8,9,0}
      fmt.Println(binarySearch(list, 1))
}