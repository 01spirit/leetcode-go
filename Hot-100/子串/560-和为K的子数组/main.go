package main

import (
	"fmt"
)

func subarraySum(nums []int, k int) int {
	n := len(nums)
	cnt := 0

	preSum := 0
	preSM := map[int]int{0: 1}
	for i := 0; i < n; i++ {
		preSum += nums[i]
		if c, ok := preSM[preSum-k]; ok {
			cnt += c
		}
		preSM[preSum]++
	}

	return cnt
}

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))
}
