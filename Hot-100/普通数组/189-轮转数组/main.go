package main

import (
	"slices"
)

func rotate(nums []int, k int) {
	k %= len(nums)
	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}

func main() {
	//fmt.Println(rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3))
	//fmt.Println(rotate([]int{-1, -100, 3, 99}, 2))
}
