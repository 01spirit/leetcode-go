package main

import (
	"fmt"
	"slices"
)

func deleteAndEarn(nums []int) int {
	m := slices.Max(nums)
	dp := make([]int, m+1)
	elesum := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := elesum[nums[i]]; ok {
			elesum[nums[i]] += nums[i]
		} else {
			elesum[nums[i]] = nums[i]
		}
	}

	if len(elesum) == 1 {
		return nums[0] * len(nums)
	}

	dp[0] = elesum[0]
	dp[1] = max(elesum[0], elesum[1])
	for i := 2; i <= m; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+elesum[i])
	}

	return dp[m]
}

func main() {
	fmt.Println(deleteAndEarn([]int{3, 4, 2}))
	fmt.Println(deleteAndEarn([]int{2, 2, 3, 3, 3, 4}))
}
