package main

import "fmt"

func sortedSquares(nums []int) []int {
	left, right := 0, len(nums)-1
	idx := len(nums) - 1
	squares := make([]int, len(nums))

	for left <= right {
		if abs(nums[left]) < abs(nums[right]) {
			squares[idx] = nums[right] * nums[right]
			right--
		} else {
			squares[idx] = nums[left] * nums[left]
			left++
		}
		idx--
	}

	return squares
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
}
