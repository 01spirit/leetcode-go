package main

import "fmt"

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	reverse := 0
	seq := 0
	left := min(start, destination)
	right := max(start, destination)
	for i := 0; i < len(distance); i++ {
		if i < left || i >= right {
			reverse += distance[i]
		} else {
			seq += distance[i]
		}

	}

	return min(seq, reverse)
}

func main() {
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 1))
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 2))
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 3))

	fmt.Println(distanceBetweenBusStops([]int{7, 10, 1, 12, 11, 14, 5, 0}, 7, 2))
}
