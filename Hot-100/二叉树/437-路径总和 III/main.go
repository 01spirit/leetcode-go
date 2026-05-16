package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	//// key：从根到 node 的节点值之和
	//// value：节点值之和的出现次数
	//// 注意在递归过程中，哈希表只保存根到 node 的路径的前缀的节点值之和
	//cnt := map[int]int{0: 1}
	//
	//var ans int
	//// s 表示从根到 node 的父节点的节点值之和（node 的节点值尚未计入）
	//var dfs func(*TreeNode, int)
	//dfs = func(node *TreeNode, s int) {
	//	if node == nil {
	//		return
	//	}
	//
	//	s += node.Val
	//	// 把 node 当作路径的终点，统计有多少个起点
	//	ans += cnt[s-targetSum]
	//
	//	cnt[s]++
	//	dfs(node.Left, s)
	//	dfs(node.Right, s)
	//	cnt[s]-- // 恢复现场（撤销 cnt[s]++）
	//}
	//
	//dfs(root, 0)
	//return ans

	ans := 0
	cnt := map[int]int{0: 1}
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		sum += root.Val
		ans += cnt[sum-targetSum]

		cnt[sum]++
		dfs(root.Left, sum)
		dfs(root.Right, sum)
		cnt[sum]--
	}
	dfs(root, 0)

	return ans
}
