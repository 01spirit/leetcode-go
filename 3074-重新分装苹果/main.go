package main

import (
	"fmt"
	"slices"
)

func minimumBoxes(apple []int, capacity []int) int {
	appleCnt := 0
	for _, num := range apple {
		appleCnt += num
	}

	slices.Sort(capacity)
	for i := len(capacity) - 1; i >= 0; i-- {
		appleCnt -= capacity[i]
		if appleCnt <= 0 {
			return len(capacity) - i
		}
	}

	return len(capacity)
}

func main() {
	fmt.Println(minimumBoxes([]int{1, 3, 2}, []int{4, 3, 1, 5, 2}))
	fmt.Println(minimumBoxes([]int{5, 5, 5}, []int{2, 4, 2, 7}))
}
