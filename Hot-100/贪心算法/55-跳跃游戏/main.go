package main

import "fmt"

func canJump(nums []int) bool {
	dst := 0
	for i := 0; i < len(nums); i++ {
		if i <= dst {
			dst = max(dst, i+nums[i])
		}
		if dst >= len(nums)-1 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(canJump([]int{2, 0, 1, 0, 1}))
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}
