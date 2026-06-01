package main

import (
	"fmt"
	"math"
)

func minElement(nums []int) int {
	ms := math.MaxInt32
	sum := 0
	for _, n := range nums {
		for n > 0 {
			x := n % 10
			n /= 10
			sum += x
		}
		ms = min(ms, sum)
		sum = 0
	}
	return ms
}

func main() {
	fmt.Println(minElement([]int{10, 12, 13, 14}))
	fmt.Println(minElement([]int{1, 2, 3, 4}))
	fmt.Println(minElement([]int{999, 19, 199}))
}
