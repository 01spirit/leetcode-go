package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var arr []int

	cur := []*TreeNode{root}
	for len(cur) > 0 {
		var tmp []*TreeNode
		for i := 0; i < len(cur); i++ {
			if cur[i].Left != nil {
				tmp = append(tmp, cur[i].Left)
			}
			if cur[i].Right != nil {
				tmp = append(tmp, cur[i].Right)
			}
			if i == len(cur)-1 {
				arr = append(arr, cur[i].Val)
			}
		}
		cur = tmp
	}

	return arr
}
