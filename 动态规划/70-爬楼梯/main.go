package main

import "fmt"

func climbStairs(n int) int {
	dp := make([]int, n+2)
	dp[0] = 1
	dp[1] = 2
	for i := 0; i < n; i++ {
		dp[i+2] = dp[i+1] + dp[i]
	}

	return dp[n-1]
}

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
}
