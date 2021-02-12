package main

import "fmt"

func main() {
	fmt.Println(NumberOfWaysToMakeChange(6, []int{1, 5}))
}

func NumberOfWaysToMakeChange(n int, denoms []int) int {
	ways := make([]int, n+1)

	ways[0] = 1

	for _, denom := range denoms {
		for amount := 1; amount < n+1; amount++ {
			if denom <= amount {
				//[0,1,2,3,4,5,6] till amount to change i.e. n
				//[1,0,0,0,0,0,0] initial ways array, is update as per formula
				//ways[1] = ways[1] + ways[1-1]
				ways[amount] += ways[amount-denom]
			}
		}
	}
	return ways[n]
}
