package main

// This is the struct of the input root. Do not edit it.
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func BranchSums(root *BinaryTree) []int {
	// Write your code here.
	sums := []int{}
	calculateBranchSums(rote,0, &sums)
	return sums
}

func calculateBranchSums(node *BinaryTree, runningSum int, *sums []int) {
	if node == nil {
		return
	}
	runningSum += node.Value
	if node.Left == nil && node.Right == nil {
		return
	}
	calculateBranchSums(node.Left, runningSums, sums)
	calculateBranchSums(node.Right, runningSums, sums)
	
}
