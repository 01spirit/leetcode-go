package main

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
	return &TreeNode{Val: val}
}

func Insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return NewNode(val)
	}

	cur := root
	for {
		if val < cur.Val {
			if cur.Left == nil {
				cur.Left = NewNode(val)
				break
			}
			cur = cur.Left
		} else if val > cur.Val {
			if cur.Right == nil {
				cur.Right = NewNode(val)
				break
			}
			cur = cur.Right
		} else {
			break
		}
	}

	return root
}

func Delete(root *TreeNode, val int) *TreeNode {
	var parent *TreeNode
	cur := root

	// 找到要删除的节点
	for cur != nil && cur.Val != val {
		parent = cur
		if val < cur.Val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	if cur == nil {
		return root
	}

	// 如果有两个孩子，找右子树最小替换
	if cur.Left != nil && cur.Right != nil {
		minParent := cur
		minNode := cur.Right
		for minNode.Left != nil {
			minParent = minNode
			minNode = minNode.Left
		}
		cur.Val = minNode.Val
		parent = minParent
		cur = minNode
	}

	// 此时 cur 至多一个孩子
	var child *TreeNode
	if cur.Left != nil {
		child = cur.Left
	} else {
		child = cur.Right
	}

	if parent == nil {
		return child // 删除的是 root
	}

	if parent.Left == cur {
		parent.Left = child
	} else {
		parent.Right = child
	}

	return root
}

func Search(root *TreeNode, val int) *TreeNode {
	cur := root
	for cur != nil {
		if val < cur.Val {
			cur = cur.Left
		} else if val > cur.Val {
			cur = cur.Right
		} else {
			return cur
		}
	}
	return nil
}

func Update(root *TreeNode, oldVal, newVal int) *TreeNode {
	node := Search(root, oldVal)
	if node == nil {
		return root
	}
	root = Delete(root, oldVal)
	root = Insert(root, newVal)
	return root
}

func PreOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	PreOrder(root.Left, res)
	PreOrder(root.Right, res)
}

func InOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	InOrder(root.Left, res)
	*res = append(*res, root.Val)
	InOrder(root.Right, res)
}

func PostOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	PostOrder(root.Left, res)
	PostOrder(root.Right, res)
	*res = append(*res, root.Val)
}
