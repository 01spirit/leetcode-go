package main

import "fmt"

func constructTransformedArray(nums []int) []int {
	length := len(nums)
	result := make([]int, length)

	for i := 0; i < length; i++ {
		result[i] = nums[((i+nums[i])%length+length)%length]
	}

	return result
}

func main() {
	fmt.Println(constructTransformedArray([]int{3, -2, 1, 1}))
	fmt.Println(constructTransformedArray([]int{-1, 4, -1}))
	fmt.Println(constructTransformedArray([]int{-10, -10}))
}
