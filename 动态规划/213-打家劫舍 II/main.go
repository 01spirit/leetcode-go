package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	n := len(nums)

	dp0 := make([]int, n)
	dp1 := make([]int, n)
	dp0[0] = nums[0]
	dp0[1] = max(nums[0], nums[1])
	dp1[1] = nums[1]
	for i := 2; i < n-1; i++ {
		dp0[i] = max(dp0[i-1], dp0[(n+i-2)%n]+nums[i])
	}
	for i := 2; i < n; i++ {
		dp1[i] = max(dp1[i-1], dp1[i-2]+nums[i])
	}

	return max(dp0[n-2], dp1[n-1])
}

func main() {
	fmt.Println(rob([]int{0, 0}))
	fmt.Println(rob([]int{1, 2, 1, 1}))
	fmt.Println(rob([]int{2, 3, 2}))
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{1, 2, 3}))
}
