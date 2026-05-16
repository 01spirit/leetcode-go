package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		return
	}
	var arr []*TreeNode
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		arr = append(arr, node)
		traverse(node.Left)
		traverse(node.Right)
	}

	traverse(root)

	for i := 0; i < len(arr)-1; i++ {
		arr[i].Left = nil
		arr[i].Right = arr[i+1]
	}
	arr[len(arr)-1].Left = nil
	arr[len(arr)-1].Right = nil
}
