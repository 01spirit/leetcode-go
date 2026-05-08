package main

import "fmt"

func countHousePlacements(n int) int {
	const mod = 1e9 + 7
	dp := make([]int, n+2)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n+1; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % mod
	}
	return (dp[n+1] * dp[n+1]) % mod
}

func main() {
	fmt.Println(countHousePlacements(3))
	fmt.Println(countHousePlacements(1))
	fmt.Println(countHousePlacements(2))
}
