package main

import (
	"fmt"
)

func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	solve := func(start1, duration1, start2, duration2 []int) int {
		finish1 := 2147483647
		for i := 0; i < len(start1); i++ {
			if val := start1[i] + duration1[i]; val < finish1 {
				finish1 = val
			}
		}
		finish2 := 2147483647
		for i := 0; i < len(start2); i++ {
			curStart := start2[i]
			if finish1 > curStart {
				curStart = finish1
			}
			if val := curStart + duration2[i]; val < finish2 {
				finish2 = val
			}
		}
		return finish2
	}

	landWater := solve(landStartTime, landDuration, waterStartTime, waterDuration)
	waterLand := solve(waterStartTime, waterDuration, landStartTime, landDuration)
	if landWater < waterLand {
		return landWater
	}
	return waterLand
}

func main() {
	fmt.Println(earliestFinishTime([]int{2, 8}, []int{4, 1}, []int{6}, []int{3}))
	fmt.Println(earliestFinishTime([]int{5}, []int{3}, []int{1}, []int{10}))
}
