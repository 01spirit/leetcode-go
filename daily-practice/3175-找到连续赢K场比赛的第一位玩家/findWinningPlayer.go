package main

import (
	"fmt"
)

/* 顺利通过，空间复杂度较高 */
//func findWinningPlayer(skills []int, k int) int {
//	mp := make(map[int]int)
//	for i, num := range skills {
//		mp[num] = i
//	}
//	if k >= len(skills) {
//		return mp[slices.Max(skills)]
//	}
//
//	cur := skills[0]
//	curCnt := 0
//	for i := 1; i < len(skills)*2; i++ {
//		if skills[i%len(skills)] < cur {
//			curCnt++
//		} else {
//			curCnt = 1
//			cur = skills[i%len(skills)]
//		}
//		if curCnt == k {
//			return mp[cur]
//		}
//	}
//
//	return mp[skills[0]]
//}

func findWinningPlayer(skills []int, k int) int {
	n := len(skills)
	cnt := 0
	left, last := 0, 0

	for left < n {
		right := left + 1
		for right < n && cnt < k && skills[right] < skills[left] {
			cnt++
			right++
		}
		if cnt == k {
			return left
		}
		last = left
		left = right
		cnt = 1
	}

	return last
}

func main() {
	fmt.Println(findWinningPlayer([]int{4, 2, 6, 3, 9}, 2))
	fmt.Println(findWinningPlayer([]int{2, 5, 4}, 3))
}
