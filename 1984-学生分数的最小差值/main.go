package main

import (
	"fmt"
	"math"
	"slices"
)

func minimumDifference(nums []int, k int) int {
	slices.Sort(nums)

	minDiff := math.MaxInt
	for i := 0; i < len(nums)-k+1; i++ {
		diff := nums[i+k-1] - nums[i]
		minDiff = min(minDiff, diff)
	}

	return minDiff
}

func main() {
	fmt.Println(minimumDifference([]int{90}, 1))
	fmt.Println(minimumDifference([]int{9, 4, 1, 7}, 2))
}
