package main

import "fmt"

func minOperations(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		if (nums[i]+ans)%2 == 0 {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(minOperations([]int{0, 0, 0, 1}))
	fmt.Println(minOperations([]int{0, 1, 1, 0, 1}))
	fmt.Println(minOperations([]int{1, 0, 0, 0}))
}
