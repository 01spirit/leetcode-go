package main

import (
	"fmt"
	"slices"
)

func smallestRangeII(nums []int, k int) int {
	slices.Sort(nums)
	ans := nums[len(nums)-1] - nums[0]
	for i := 0; i < len(nums)-1; i++ {
		ans = min(ans, polar(nums, k, i))
	}
	return ans
}

func polar(nums []int, k, index int) int {
	ax := max(nums[index]+k, nums[len(nums)-1]-k)
	in := min(nums[0]+k, nums[index+1]-k)
	return ax - in
}

func main() {
	fmt.Println(smallestRangeII([]int{1}, 0))
	fmt.Println(smallestRangeII([]int{0, 10}, 2))
	fmt.Println(smallestRangeII([]int{1, 3, 6}, 3))
}
