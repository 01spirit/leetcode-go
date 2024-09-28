package main

import "fmt"

func isArraySpecial(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	odd := nums[0] % 2
	for i, num := range nums {
		if i%2 == 0 && num%2 != odd {
			return false
		} else if i%2 == 1 && num%2 == odd {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isArraySpecial([]int{1}))
	fmt.Println(isArraySpecial([]int{2, 1, 4}))
	fmt.Println(isArraySpecial([]int{4, 3, 1, 6}))
}
