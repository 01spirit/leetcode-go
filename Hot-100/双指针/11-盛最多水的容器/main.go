package main

import (
	"fmt"
)

func maxArea(height []int) int {
	ans := 0
	left, right := 0, len(height)-1
	for left < right {
		ans = max(ans, (right-left)*min(height[left], height[right]))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

func main() {
	fmt.Println(maxArea([]int{1, 2, 4, 3}))
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
}
