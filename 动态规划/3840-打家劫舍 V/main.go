package main

import "fmt"

func subarray(nums []int, colors []int) int64 {
	n := len(nums)
	if n == 1 {
		return int64(nums[0])
	}
	dp := make([]int64, n)

	dp[0] = int64(nums[0])
	dp[1] = int64(max(nums[0], nums[1]))
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+int64(nums[i]))
	}

	return dp[n-1]
}

func rob(nums []int, colors []int) int64 {
	start := 0
	ans := int64(0)

	for i := 0; i < len(nums)-1; i++ {
		if colors[i+1] != colors[i] {
			ans += subarray(nums[start:i+1], colors[start:i+1])
			start = i + 1
		}
	}
	ans += subarray(nums[start:], colors[start:])

	return ans
}

func main() {
	fmt.Println(rob([]int{1, 4, 3, 5}, []int{1, 1, 2, 2}))
	fmt.Println(rob([]int{3, 1, 2, 4}, []int{2, 3, 2, 2}))
	fmt.Println(rob([]int{10, 1, 3, 9}, []int{1, 1, 1, 2}))
}
