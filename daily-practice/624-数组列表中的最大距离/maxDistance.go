package main

import (
	"fmt"
	"math"
)

func maxDistance(arrays [][]int) int {
	minVal := math.MaxInt32
	maxVal := math.MinInt32

	sminVal := math.MaxInt32
	smaxVal := math.MinInt32

	iidx := -1
	aidx := -1
	for i, array := range arrays {
		if array[0] < minVal {
			sminVal = min(minVal, sminVal)
			minVal = array[0]
			iidx = i
		}
		if array[len(array)-1] >= maxVal {
			smaxVal = max(maxVal, smaxVal)
			maxVal = array[len(array)-1]
			aidx = i
		}
	}

	if iidx != aidx {
		return maxVal - minVal
	}

	return maxVal - minVal - min(sminVal-minVal, maxVal-smaxVal)
}

func main() {
	fmt.Println(maxDistance([][]int{{1, 4}, {0, 5}}))
	fmt.Println(maxDistance([][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}}))
	fmt.Println(maxDistance([][]int{{1}, {1}}))
}
