package main

import (
	"fmt"
)

func maxResult(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	queue := make([]int, n) // 模拟双端队列
	qi, qj := 0, 1
	for i := 1; i < n; i++ {
		for qi < qj && queue[qi] < i-k {
			qi++
		}
		dp[i] = dp[queue[qi]] + nums[i]
		for qi < qj && dp[queue[qj-1]] <= dp[i] {
			qj--
		}
		queue[qj] = i
		qj++
	}
	return dp[n-1]
}

func main() {
	//nums := []int{1, -1, -2, 4, -7, 3}
	//k := 2
	//nums := []int{10, -5, -2, 4, 0, 3}
	//k := 3
	nums := []int{1, -5, -20, 4, -1, 3, -6, -3}
	k := 2
	res := maxResult(nums, k)
	fmt.Println(res)
}
