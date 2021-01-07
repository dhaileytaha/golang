package main

//BinaryTree...
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func nodeDepths(root *BinaryTree) int {
	nodeDepthsHelper(root, 0)
}

func nodeDepthsHelper(root, depth int) int {
	if root == nil {
		return 0
	}
	return depth + nodeDepthsHelper(root.Left, depth+1) + nodeDepthsHelper(root.Right, depth+1)
}
