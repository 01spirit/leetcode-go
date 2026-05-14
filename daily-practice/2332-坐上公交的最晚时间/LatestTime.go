package main

import (
	"fmt"
	"sort"
)

func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
	sort.Ints(buses)
	sort.Ints(passengers)
	pos := 0
	space := 0

	for _, arrive := range buses {
		space = capacity
		for space > 0 && pos < len(passengers) && passengers[pos] <= arrive {
			space--
			pos++
		}
	}

	pos--
	lastCatchTime := buses[len(buses)-1]
	if space <= 0 {
		lastCatchTime = passengers[pos]
	}
	for pos >= 0 && passengers[pos] == lastCatchTime {
		pos--
		lastCatchTime--
	}

	return lastCatchTime
}

func main() {
	fmt.Println(latestTimeCatchTheBus([]int{3, 2}, []int{2, 3}, 2))
	fmt.Println(latestTimeCatchTheBus([]int{3}, []int{2, 3}, 2))
	fmt.Println(latestTimeCatchTheBus([]int{10, 20}, []int{2, 17, 18, 19}, 2))
	fmt.Println(latestTimeCatchTheBus([]int{20, 30, 10}, []int{19, 13, 26, 4, 25, 11, 21}, 2))
}
