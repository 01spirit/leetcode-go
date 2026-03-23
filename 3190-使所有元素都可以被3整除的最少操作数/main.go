package main

import "fmt"

func minimumOperations(nums []int) int {
	ans := 0

	for _, num := range nums {
		ans += min(num%3, 3-num%3)
	}

	return ans
}

func main() {
	fmt.Println(minimumOperations([]int{1, 2, 3, 4}))
	fmt.Println(minimumOperations([]int{3, 6, 9}))
}
