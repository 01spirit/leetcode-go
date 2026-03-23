package main

import "fmt"

func prefixesDivBy5(nums []int) []bool {
	ans := make([]bool, len(nums))

	val := 0
	for i := 0; i < len(nums); i++ {
		val = ((val << 1) | nums[i]) % 5
		ans[i] = val == 0
	}

	return ans
}

func main() {
	fmt.Println(prefixesDivBy5([]int{1, 0, 1}))
	fmt.Println(prefixesDivBy5([]int{0, 1, 1}))
	fmt.Println(prefixesDivBy5([]int{1, 1, 1}))
}
