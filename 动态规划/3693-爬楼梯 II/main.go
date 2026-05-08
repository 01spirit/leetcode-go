package main

import "fmt"

func climbStairs(n int, costs []int) int {
	dp := make([]int, n+1)
	if len(costs) == 1 {
		return costs[0] + 1
	}
	if len(costs) == 2 {
		return min(costs[0]+1+costs[1]+1, costs[1]+4)
	}

	dp[0] = 0
	dp[1] = costs[0] + 1
	dp[2] = min(dp[1]+costs[1]+1, dp[0]+costs[1]+4)
	for i := 3; i <= n; i++ {
		dp[i] = min(dp[i-1]+costs[i-1]+1, dp[i-2]+costs[i-1]+4, dp[i-3]+costs[i-1]+9)
	}

	return dp[n]
}

func main() {
	fmt.Println(climbStairs(4, []int{1, 2, 3, 4}))
	fmt.Println(climbStairs(4, []int{5, 1, 6, 2}))
	fmt.Println(climbStairs(3, []int{9, 8, 3}))
}
