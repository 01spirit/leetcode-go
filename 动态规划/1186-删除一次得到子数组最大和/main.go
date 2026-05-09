package main

import (
	"fmt"
	"math"
)

func maximumSum(arr []int) int {
	ans := math.MinInt / 2
	f0, f1 := ans, ans
	for _, x := range arr {
		f1 = max(f1+x, f0)
		f0 = max(f0, 0) + x
		ans = max(ans, f0, f1)
	}
	return ans
}

func main() {
	fmt.Println(maximumSum([]int{-1, -2}))
	fmt.Println(maximumSum([]int{2, 1, -2, -5, -2}))
	fmt.Println(maximumSum([]int{-7, 6, 1, 2, 1, 4, -1}))
	fmt.Println(maximumSum([]int{1, -2, 0, 3}))
	fmt.Println(maximumSum([]int{1, -2, -2, 3}))
	fmt.Println(maximumSum([]int{-1, -1, -1, -1}))
}
