package main

import (
	"fmt"
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	n := len(nums)
	bitsMaxPos := make([]int, 31)
	for i := range bitsMaxPos {
		bitsMaxPos[i] = -1
	}

	res := math.MaxInt
	for i := 0; i < n; i++ {
		for j := 0; j <= 30; j++ {
			if nums[i]>>j&1 == 1 {
				bitsMaxPos[j] = i
			}
		}

		posToBit := make([][2]int, 0)
		for j := 0; j <= 30; j++ {
			if bitsMaxPos[j] != -1 {
				posToBit = append(posToBit, [2]int{bitsMaxPos[j], j})
			}
		}
		sort.Slice(posToBit, func(a, b int) bool {
			return posToBit[a][0] > posToBit[b][0]
		})

		val := 0
		for j, p := 0, 0; j < len(posToBit); p = j {
			for j < len(posToBit) && posToBit[j][0] == posToBit[p][0] {
				val |= 1 << posToBit[j][1]
				j++
			}
			res = min(res, int(math.Abs(float64(val-k))))
		}
	}
	return res
}

func main() {
	fmt.Println(minimumDifference([]int{10, 75, 60, 36, 17}, 54)) // 1
	fmt.Println(minimumDifference([]int{8, 89, 23, 12, 13}, 59))  // 28
	fmt.Println(minimumDifference([]int{3, 50, 1, 29, 27}, 66))   // 3
	fmt.Println(minimumDifference([]int{2, 48, 34, 28, 52}, 62))  // 0
	fmt.Println(minimumDifference([]int{1, 8, 9}, 7))             // 1
	fmt.Println(minimumDifference([]int{1, 2, 4, 5}, 3))          // 0
	fmt.Println(minimumDifference([]int{1, 3, 1, 3}, 2))          // 1
	fmt.Println(minimumDifference([]int{1}, 10))                  // 9
}
