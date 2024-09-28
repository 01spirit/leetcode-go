package main

import "fmt"

func isArraySpecial(nums []int, queries [][]int) []bool {
	ans := make([]bool, len(queries))
	dp := make([]int, len(nums))

	for i := range nums {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		if (nums[i]^nums[i-1])&1 == 1 { // 奇偶性不同
			dp[i] = dp[i-1] + 1
		}
	}

	for i, query := range queries {
		x, y := query[0], query[1]
		ans[i] = dp[y] >= y-x+1
	}

	return ans
}

func main() {
	fmt.Println(isArraySpecial([]int{3, 4, 1, 2, 6}, [][]int{{0, 4}}))
	fmt.Println(isArraySpecial([]int{4, 3, 1, 6}, [][]int{{0, 2}, {2, 3}}))
}
