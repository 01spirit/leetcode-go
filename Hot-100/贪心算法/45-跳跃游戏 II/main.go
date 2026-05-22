package main

import "fmt"

//func jump(nums []int) int {
//	if len(nums) == 1 {
//		return 0
//	}
//	cnt := 0
//	i := 0
//	for i < len(nums)-1 {
//		md := 0
//		mp := i
//		for j := i + 1; j <= i+nums[i] && j < len(nums); j++ {
//			if j+nums[j] >= md || j+nums[j] >= len(nums)-1 {
//				md = j + nums[j]
//				mp = j
//			}
//		}
//		i = mp
//		cnt++
//	}
//	return cnt
//}

func jump(nums []int) int {
	cnt := 0
	maxPos := 0
	end := 0
	for i := 0; i < len(nums)-1; i++ {
		maxPos = max(maxPos, i+nums[i])
		if i == end {
			end = maxPos
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(jump([]int{2, 3, 1}))
	fmt.Println(jump([]int{1, 2}))
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump([]int{2, 3, 0, 1, 4}))
}
