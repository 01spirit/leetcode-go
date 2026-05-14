package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var ans [][]int

	cur := []*TreeNode{root}
	for len(cur) > 0 {
		n := len(cur)
		arr := make([]int, 0)
		var level []*TreeNode
		for i := 0; i < n; i++ {
			arr = append(arr, cur[i].Val)
			if cur[i].Left != nil {
				level = append(level, cur[i].Left)
			}
			if cur[i].Right != nil {
				level = append(level, cur[i].Right)
			}
		}
		cur = level
		ans = append(ans, arr)
	}

	return ans
}

func main() {
	fmt.Println(levelOrder(&TreeNode{Val: 1, Left: nil, Right: nil}))
}
