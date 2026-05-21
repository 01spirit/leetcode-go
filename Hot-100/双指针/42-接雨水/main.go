package main

import "fmt"

/*
	前后缀分离
*/
//func trap(height []int) int {
//	n := len(height)
//	pre := make([]int, n)
//	suf := make([]int, n)
//
//	pre[0] = height[0]
//	for i := 1; i < n; i++ {
//		pre[i] = max(pre[i-1], height[i])
//	}
//	suf[n-1] = height[n-1]
//	for i := n - 2; i >= 0; i-- {
//		suf[i] = max(suf[i+1], height[i])
//	}
//
//	ans := 0
//	for i := 1; i < n; i++ {
//		ans += min(pre[i], suf[i]) - height[i]
//	}
//
//	return ans
//}

/*
双指针
*/
func trap(height []int) int {
	n := len(height)
	preMax, sufMax := 0, 0
	left, right := 0, n-1
	ans := 0
	for left < right {
		preMax = max(preMax, height[left])
		sufMax = max(sufMax, height[right])
		if height[left] < height[right] {
			ans += preMax - height[left]
			left++
		} else {
			ans += sufMax - height[right]
			right--
		}
	}
	return ans
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
}
