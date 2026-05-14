package main

import (
	"slices"
)

func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	res := make([][]int, 0)
	for first := 0; first < len(nums); first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := len(nums) - 1
		target := -nums[first]

		for second := first + 1; second < len(nums); second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}

	res := threeSum(nums)

	for _, num := range res {
		for _, n := range num {
			print(n)
		}
		println()
	}
}
