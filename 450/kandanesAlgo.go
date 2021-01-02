package main

func KadanesAlgorithm(array []int) int {
	// Write your code here.
	maxEndingHere := array[0]
	maxSoFar := array[0]
	for i := 1; i < len(array); i++ {
		num := array[i]
		maxEndingHere = max(num, maxEndingHere+num)
		maxSoFar = max(maxSoFar, maxEndingHere)
	}
	return maxSoFar
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
