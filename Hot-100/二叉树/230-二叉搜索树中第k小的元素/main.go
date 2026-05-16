package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func kthSmallest(root *TreeNode, k int) int {
//	var arr []int
//	traverse(root, &arr)
//	slices.Sort(arr)
//	return arr[k-1]
//}
//
//func traverse(root *TreeNode, arr *[]int) {
//	if root == nil {
//		return
//	}
//	*arr = append(*arr, root.Val)
//	traverse(root.Left, arr)
//	traverse(root.Right, arr)
//}

func kthSmallest(root *TreeNode, k int) int {
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		leftRes := dfs(node.Left)
		if leftRes != -1 {
			return leftRes
		}
		k--
		if k == 0 {
			leftRes = node.Val
			return leftRes
		}
		leftRes = dfs(node.Right)
		return leftRes
	}
	return dfs(root)
}
