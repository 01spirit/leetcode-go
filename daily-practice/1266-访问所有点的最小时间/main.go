package main

import (
	"fmt"
	"math"
)

func minTimeToVisitAllPoints(points [][]int) int {
	ans := 0

	for i := 1; i < len(points); i++ {
		x0 := points[i-1][0]
		y0 := points[i-1][1]
		x1 := points[i][0]
		y1 := points[i][1]

		z := int(min(math.Abs(float64(x1-x0)), math.Abs(float64(y1-y0))))

		ans += z
		ans += int(max(math.Abs(float64(x1-x0)), math.Abs(float64(y1-y0)))) - z
	}

	return ans
}

func main() {
	fmt.Println(minTimeToVisitAllPoints([][]int{[]int{1, 1}, []int{3, 4}, []int{-1, 0}}))
	fmt.Println(minTimeToVisitAllPoints([][]int{[]int{3, 2}, []int{-2, 2}}))
}
