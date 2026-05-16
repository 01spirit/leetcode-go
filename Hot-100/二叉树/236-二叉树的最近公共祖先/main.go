package main

import (
	"fmt"
	"math"
)

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

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var arr [][]int
	var search func(node *TreeNode, val int, path []int)
	search = func(node *TreeNode, val int, path []int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		if node.Val == val {
			arr = append(arr, append([]int(nil), path...))
			return
		}
		search(node.Left, val, path)
		search(node.Right, val, path)
	}
	search(root, p.Val, []int{})
	search(root, q.Val, []int{})

	var find func(node *TreeNode, val int) *TreeNode
	find = func(node *TreeNode, val int) *TreeNode {
		if node == nil {
			return nil
		}
		if node.Val == val {
			return node
		}
		var ret *TreeNode
		ret = find(node.Left, val)
		if ret != nil {
			return ret
		}
		ret = find(node.Right, val)
		return ret
	}
	for i := 0; i < len(arr[0]) && i < len(arr[1]); i++ {
		if arr[0][i] == arr[1][i] {
			continue
		} else {
			return find(root, arr[0][i-1])
		}
	}

	if len(arr[0]) < len(arr[1]) {
		return find(root, arr[0][len(arr[0])-1])
	}

	return find(root, arr[1][len(arr[1])-1])
}

func main() {
	fmt.Println(lowestCommonAncestor(BuildTree([]int{3, 5, 1, 6, 2, 0, 8, math.MinInt32, math.MinInt32, 7, 4}), BuildTree([]int{5}), BuildTree([]int{1})))
}
