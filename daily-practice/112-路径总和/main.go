package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	find := false
	var dfs func(root *TreeNode, sum int)
	dfs = func(node *TreeNode, sum int) {
		if find {
			return
		}
		if node == nil {
			return
		}
		sum += node.Val
		if sum == targetSum && node.Left == nil && node.Right == nil {
			find = true
			return
		}
		dfs(node.Left, sum)
		dfs(node.Right, sum)
	}
	dfs(root, 0)
	return find
}
