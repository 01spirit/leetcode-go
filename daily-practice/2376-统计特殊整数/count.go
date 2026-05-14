package main

import "fmt"

// 我觉得没问题，但是超出内存限制了
func countSpecialNumbers(n int) int {
	ans := 0

	dp := make([]int16, n+1)
	dpb := make([]byte, n+1)

	dp[0] = 0
	dpb[0] = 0

	for i := 1; i <= n; i++ {
		remain := i / 10
		low := i % 10
		if dpb[remain] == 1 || dp[remain]&(1<<low) != 0 {
			dpb[i] = 1
		} else {
			ans++
		}
		dp[i] = dp[remain]
		dp[i] |= 1 << low
	}

	return ans
}

func main() {
	fmt.Println(countSpecialNumbers(20))
	fmt.Println(countSpecialNumbers(5))
	fmt.Println(countSpecialNumbers(135))
}
