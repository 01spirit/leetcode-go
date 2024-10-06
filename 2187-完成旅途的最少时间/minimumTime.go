package main

import (
	"fmt"
	"slices"
)

func minimumTime(time []int, totalTrips int) int64 {

	left, right := int64(1), int64(totalTrips)*int64(slices.Min(time))
	for left < right {
		mid := int64(left + (right-left)/2)
		trips := int64(0)
		for i := range time {
			trips += mid / int64(time[i])
		}
		if trips < int64(totalTrips) {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return right
}

func main() {
	fmt.Println(minimumTime([]int{1, 2, 3}, 5))
	fmt.Println(minimumTime([]int{2}, 1))
}
