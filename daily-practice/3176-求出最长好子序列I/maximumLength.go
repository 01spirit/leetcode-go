package main

func maximumLength(nums []int, k int) int {
	result := 0

	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, 51)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	for i := range nums {
		dp[i][0] = 1
		for l := 0; l <= k; l++ {
			for j := 0; j < i; j++ {
				add := 0
				if nums[i] != nums[j] {
					add = 1
				}

				if l-add >= 0 && dp[j][l-add] != -1 {
					dp[i][l] = max(dp[i][l], dp[j][l-add]+1)
				}

			}
			result = max(result, dp[i][l])
		}
	}

	return result
}
