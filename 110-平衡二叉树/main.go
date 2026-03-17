package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return checkHeight(root) != -1
}

func checkHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftHeight := checkHeight(node.Left)
	if leftHeight == -1 {
		return -1
	}

	rightHeight := checkHeight(node.Right)
	if rightHeight == -1 {
		return -1
	}

	if math.Abs(float64(leftHeight-rightHeight)) > 1 {
		return -1
	}

	return max(leftHeight, rightHeight) + 1
}

func main() {

}
