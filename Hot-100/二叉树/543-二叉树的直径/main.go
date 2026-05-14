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

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	deepen(root, &ans)
	return ans
}

func deepen(root *TreeNode, ans *int) (int, int) {
	if root == nil {
		return 0, 0
	}
	left, right := 0, 0
	if root.Left != nil {
		left = 1
	}
	if root.Right != nil {
		right = 1
	}
	left += max(deepen(root.Left, ans))
	right += max(deepen(root.Right, ans))
	*ans = max(*ans, left+right)
	return left, right
}
