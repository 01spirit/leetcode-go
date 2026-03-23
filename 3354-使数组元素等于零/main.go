package main

import (
	"fmt"
)

func countValidSelections(nums []int) int {
	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}

	ans, leftSum, rightSum := 0, 0, sum
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			if leftSum-rightSum >= 0 && leftSum-rightSum <= 1 {
				ans++
			}
			if rightSum-leftSum >= 0 && rightSum-leftSum <= 1 {
				ans++
			}
		} else {
			leftSum += nums[i]
			rightSum -= nums[i]
		}
	}

	return ans
}

func main() {
	fmt.Println(countValidSelections([]int{1, 0, 2, 0, 3}))
	fmt.Println(countValidSelections([]int{2, 3, 4, 0, 4, 1, 0}))
}
