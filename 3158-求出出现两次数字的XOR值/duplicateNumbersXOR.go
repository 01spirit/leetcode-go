package main

import "fmt"

func duplicateNumbersXOR(nums []int) int {
	ans := 0

	cnt := int64(0)
	for _, num := range nums {
		if cnt&(1<<num) > 0 {
			ans ^= num
		} else {
			cnt |= 1 << num
		}
	}

	return ans
}

func main() {
	fmt.Println(duplicateNumbersXOR([]int{1, 2, 1, 3}))
	fmt.Println(duplicateNumbersXOR([]int{1, 2, 3}))
	fmt.Println(duplicateNumbersXOR([]int{1, 2, 2, 1}))
}
