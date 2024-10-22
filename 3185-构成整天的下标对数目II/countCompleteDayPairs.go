package main

import "fmt"

func countCompleteDayPairs(hours []int) int64 {
	diffBucket := make([]map[int]int64, 0)
	for i := 0; i < 6; i++ {
		diffBucket = append(diffBucket, make(map[int]int64))
	}
	ans := int64(0)

	for _, num := range hours {
		ans += diffBucket[num%6][num%24]
		diffBucket[(24-num%24)%6][(24-num%24)%24]++
	}

	return ans
}

func main() {
	fmt.Println(countCompleteDayPairs([]int{12, 12, 30, 24, 24}))
	fmt.Println(countCompleteDayPairs([]int{72, 48, 24, 3}))
}
