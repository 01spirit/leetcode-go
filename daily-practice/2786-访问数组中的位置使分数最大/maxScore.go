package main

import (
	"fmt"
	"math"
)

func maxScore(nums []int, x int) int64 {
	dp := [2]int{math.MinInt32, math.MinInt32}
	res := int64(nums[0])
	dp[nums[0]%2] = nums[0]
	for i, num := range nums {
		if i == 0 {
			continue
		}
		parity := num % 2
		cur := max(dp[parity]+num, dp[1-parity]+num-x)
		res = max(res, int64(cur))
		dp[parity] = max(dp[parity], cur)
	}
	return res
}

func main() {
	fmt.Println(maxScore([]int{2, 3, 6, 1, 9, 2}, 5))
	fmt.Println(maxScore([]int{2, 4, 6, 8}, 3))
}
