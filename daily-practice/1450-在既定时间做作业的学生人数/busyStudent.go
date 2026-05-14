package main

import "fmt"

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	var count int = 0

	for i := 0; i < len(startTime); i++ {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			count++
		}
	}

	return count
}

func main() {
	startTime := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	endTime := []int{10, 10, 10, 10, 10, 10, 10, 10, 10}
	queryTime := 50

	count := busyStudent(startTime, endTime, queryTime)

	fmt.Println(count)
}
