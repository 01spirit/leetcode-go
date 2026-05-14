package main

import "fmt"

func numberOfPoints(nums [][]int) int {

	high := 0
	for _, arr := range nums {
		high = max(high, arr[1])
	}

	count := make([]int, high)
	for _, arr := range nums {
		for i := arr[0]; i <= arr[1]; i++ {
			count[i-1]++
		}
	}

	result := 0
	for _, num := range count {
		if num > 0 {
			result++
		}
	}

	return result
}

func main() {
	fmt.Println(numberOfPoints([][]int{{3, 6}, {1, 5}, {4, 7}}))
	fmt.Println(numberOfPoints([][]int{{1, 3}, {5, 8}}))
}
