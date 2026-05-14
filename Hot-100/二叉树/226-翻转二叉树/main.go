package main

import "math"

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

func invertTree(root *TreeNode) *TreeNode {
	return invert(root)
}

func invert(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invert(root.Left)
	invert(root.Right)
	return root
}

func main() {

}
