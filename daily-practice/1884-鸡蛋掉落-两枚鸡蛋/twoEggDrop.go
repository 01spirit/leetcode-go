package main

import (
	"fmt"
	"math"
)

func twoEggDrop(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt32 / 2
	}
	dp[0] = 0

	for i := 1; i <= n; i++ {
		for k := 1; k <= i; k++ {
			dp[i] = min(max(k-1, dp[i-k])+1, dp[i])
		}
	}

	return dp[n]
}

func main() {
	fmt.Println(twoEggDrop(2))
	fmt.Println(twoEggDrop(6))
	fmt.Println(twoEggDrop(100))
}
