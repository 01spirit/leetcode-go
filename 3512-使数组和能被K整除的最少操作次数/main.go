package main

import "fmt"

func minOperations(nums []int, k int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum % k
}

func main() {
	fmt.Println(minOperations([]int{3, 9, 7}, 5))
	fmt.Println(minOperations([]int{4, 1, 3}, 4))
	fmt.Println(minOperations([]int{3, 2}, 6))
}
