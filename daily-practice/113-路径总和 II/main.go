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

func pathSum(root *TreeNode, targetSum int) [][]int {
	var ans [][]int
	var path []int
	var dfs func(node *TreeNode, sum int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		sum += node.Val
		path = append(path, node.Val)
		defer func() {
			path = path[:len(path)-1]
		}()
		if sum == targetSum && node.Left == nil && node.Right == nil {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		dfs(node.Left, sum)
		dfs(node.Right, sum)
	}
	dfs(root, 0)
	return ans
}

func main() {
	fmt.Println(pathSum(BuildTree([]int{5, 4, 8, 11, math.MinInt32, 13, 4, 7, 2, math.MinInt32, math.MinInt32, 5, 1}), 22))
}
