package main

import "fmt"

func main() {
	arr := []int{75, 105, 120, 75, 90, 135}
	fmt.Println(maxSubsetSumNoAdjacent(arr))
}

func maxSubsetSumNoAdjacent(array []int) int {
	if len(array) == 0 {
		return 0
	} else if len(array) == 1 {
		return array[0]
	}

	//create your own array (slice)
	maxSums := make([]int, len(array))
	//slice content going to be the max sum of i-2 (no adjacent) and i
	//if the above sum is less than current index i then use current i

	//starting point at 0 and 1 need to be defined
	//because we are caclulation sum of i-2 (no adjacent) and i
	maxSums[0], maxSums[1] = array[0], max(array[0], array[1])

	// the formula is maxSum = max(
	// because if previous is larger then we go for it
	//maxSums[i-1],
	//maxSum[i-2] + array[i]
	//)
	//this formula is v imp

	for i := 2; i < len(array); i++ {
		maxSums[i] = max(maxSums[i-1], maxSums[i-2]+array[i])
	}

	return maxSums[len(maxSums)-1]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
