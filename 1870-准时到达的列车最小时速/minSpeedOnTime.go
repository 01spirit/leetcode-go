package main

import (
	"fmt"
	"math"
)

func minSpeedOnTime(dist []int, hour float64) int {
	hourInt := int(math.Round(hour * 100))
	if hourInt <= (len(dist)-1)*100 {
		return -1
	}

	left, right := 1, 10000000
	for left < right {
		mid := left + (right-left)/2
		took := 0
		for i := 0; i < len(dist)-1; i++ {
			took += (dist[i]-1)/mid + 1 // 向上取整
		}
		took = took*100*mid + dist[len(dist)-1]*100

		if took <= hourInt*mid {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}

func main() {
	fmt.Println(minSpeedOnTime([]int{1, 3, 2}, 6))
	fmt.Println(minSpeedOnTime([]int{1, 3, 2}, 2.7))
	fmt.Println(minSpeedOnTime([]int{1, 3, 2}, 1.9))
}
