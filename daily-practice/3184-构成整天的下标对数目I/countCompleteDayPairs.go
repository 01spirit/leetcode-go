package main

import "fmt"

func countCompleteDayPairs(hours []int) int {
	diff := make(map[int]int)
	ans := 0

	for _, num := range hours {
		ans += diff[num%24]
		diff[(24-num%24)%24]++
	}

	return ans
}

func main() {
	fmt.Println(countCompleteDayPairs([]int{12, 12, 30, 24, 24}))
	fmt.Println(countCompleteDayPairs([]int{72, 48, 24, 3}))
}
