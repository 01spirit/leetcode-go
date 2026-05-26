package main

import (
	"fmt"
	"math"
)

func getMinDistance(nums []int, target int, start int) int {
	minDist := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			minDist = min(minDist, abs(i-start))
			if minDist == 0 {
				return minDist
			}
		}
	}
	return minDist
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(getMinDistance([]int{1, 2, 3, 4, 5}, 5, 3))
	fmt.Println(getMinDistance([]int{1}, 1, 0))
	fmt.Println(getMinDistance([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1, 0))
}
