package main

import "fmt"

func singleNumber(nums []int) int {
	single := 0
	for _, n := range nums {
		single ^= n
	}
	return single
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(singleNumber([]int{1}))
}
