package main

import (
	"fmt"
	"math"
)

func maxAbsoluteSum(nums []int) int {
	ans := 0
	preSum := 0
	minPreSum := 0
	for i := 0; i < len(nums); i++ {
		preSum += nums[i]
		ans = max(ans, preSum-minPreSum)
		minPreSum = min(minPreSum, preSum)
	}

	ans2 := 0
	preSum = 0
	maxPreSum := 0
	for i := 0; i < len(nums); i++ {
		preSum += nums[i]
		ans2 = min(ans2, preSum-maxPreSum)
		maxPreSum = max(maxPreSum, preSum)
	}

	return max(int(math.Abs(float64(ans2))), ans)
}

func main() {
	fmt.Println(maxAbsoluteSum([]int{1, -3, 2, 3, -4}))
	fmt.Println(maxAbsoluteSum([]int{2, -5, 1, -4, 3, -2}))
}
