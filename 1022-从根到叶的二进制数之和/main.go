package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, val int) int {
	if node == nil {
		return 0
	}

	val = (val << 1) | node.Val
	if node.Left == nil && node.Right == nil {
		return val
	}

	return dfs(node.Left, val) + dfs(node.Right, val)
}

func sumRootToLeaf(root *TreeNode) int {
	var sum int = 0

	sum = dfs(root, 0)

	return sum
}

func main() {

}
