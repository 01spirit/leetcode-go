package main

import (
	"fmt"
	"math"
)

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
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func Insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return NewNode(val)
	}
	//cur := root
	//for {
	//	if val < cur.Val {
	//		if cur.Left == nil {
	//			cur.Left = NewNode(val)
	//			break
	//		}
	//		cur = cur.Left
	//	} else if val > cur.Val {
	//		if cur.Right == nil {
	//			cur.Right = NewNode(val)
	//			break
	//		}
	//		cur = cur.Right
	//	} else {
	//		break
	//	}
	//}
	if val < root.Val {
		root.Left = Insert(root.Left, val)
	} else if val > root.Val {
		root.Right = Insert(root.Right, val)
	}
	return root
}

func Build(arr []int) *TreeNode {
	var root *TreeNode
	for _, v := range arr {
		if v == math.MinInt32 {
			continue
		}
		root = Insert(root, v)
	}
	return root
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	inOrder(root, &res)
	return res
}

func inOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, res)
	*res = append(*res, root.Val)
	inOrder(root.Right, res)
}

func main() {
	fmt.Println(inorderTraversal(Build([]int{1, math.MinInt32, 2, 3})))
	fmt.Println(inorderTraversal(Build([]int{})))
	fmt.Println(inorderTraversal(Build([]int{1})))
}
