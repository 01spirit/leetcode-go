package main

import "math"

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

func NewNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func Insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return NewNode(val)
	}
	if val < root.Val {
		root.Left = Insert(root.Left, val)
	} else if val > root.Val {
		root.Right = Insert(root.Right, val)
	}
	return root
}

func BuildTree(arr []int) *TreeNode {
	var root *TreeNode
	for _, num := range arr {
		if num == math.MinInt32 {
			continue
		}
		root = Insert(root, num)
	}
	return root
}

func maxDepth(root *TreeNode) int {
	return dfs(root)
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 1
	depth += max(dfs(root.Left), dfs(root.Right))
	return depth
}

func main() {

}
