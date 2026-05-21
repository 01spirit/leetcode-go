package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	ans := make([][]int, 0)
	ans = append(ans, intervals[0])
	n := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= ans[n][1] {
			ans[n][1] = max(ans[n][1], intervals[i][1])
		} else {
			ans = append(ans, intervals[i])
			n++
		}
	}

	return ans
}

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
	fmt.Println(merge([][]int{{4, 7}, {1, 4}}))
}
