package main

import "fmt"

func countGoodStrings(low int, high int, zero int, one int) int {
	const mod = 1e9 + 7
	dp := make([]int, high+1)
	dp[0] = 1
	ans := 0
	for i := 1; i <= high; i++ {
		if i >= zero {
			dp[i] = dp[i-zero]
		}
		if i >= one {
			dp[i] = (dp[i] + dp[i-one]) % mod
		}
		if i >= low {
			ans = (ans + dp[i]) % mod
		}
	}

	return ans
}

func main() {
	fmt.Println(countGoodStrings(5, 5, 2, 4))
	fmt.Println(countGoodStrings(3, 3, 1, 1))
	fmt.Println(countGoodStrings(2, 3, 1, 2))
}
