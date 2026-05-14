package main

import (
	"fmt"
	"slices"
)

func maxStrength(nums []int) int64 {
	if len(nums) == 1 {
		return int64(nums[0])
	}
	slices.Sort(nums)

	var result int64 = 0
	if nums[len(nums)-1] > 0 {
		result = 1
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] <= 0 {
			break
		}
		result *= int64(nums[i])
	}

	minus := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 {
			break
		}
		minus++
		if minus == 2 && result == 0 {
			result = 1
		}
		if minus%2 == 0 {
			result *= int64(nums[i] * nums[i-1])
		}
	}

	return result
}

func main() {
	fmt.Println(maxStrength([]int{3, -1, -5, 2, 5, -9}))
	fmt.Println(maxStrength([]int{-4, -5, -4}))
	fmt.Println(maxStrength([]int{9, 6, 3}))
}
