package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := root.Val

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftSum := max(dfs(node.Left), 0)
		rightSum := max(dfs(node.Right), 0)
		maxSum = max(maxSum, leftSum+rightSum+node.Val)

		return max(leftSum, rightSum) + node.Val
	}
	dfs(root)

	return maxSum
}
