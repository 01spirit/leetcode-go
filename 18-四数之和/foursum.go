package main

import (
	"math"
	"slices"
)

// 没做出来
func fourSum(nums []int, target int) [][]int {
	slices.Sort(nums)
	res := make([][]int, 0)
	n := len(nums)

	first := 0
	last := n - 1
	for first < last {
		//for first > 0 && first < len(nums) && nums[first] == nums[first-1] {
		//	first++
		//}
		//for last >= 0 && last < n-1 && nums[last] == nums[last+1] {
		//	last--
		//}

		left := first + 1
		right := last - 1
		maxVal := math.MinInt32
		minVal := math.MaxInt32
		for left < right {
			//for left > first+1 && nums[left] == nums[left-1] {
			//	left++
			//}
			//
			//for right < last-1 && nums[right] == nums[right+1] {
			//	right--
			//}

			sum := nums[first] + nums[last] + nums[left] + nums[right]
			maxVal = max(sum, maxVal)
			minVal = min(sum, minVal)
			if sum == target {
				res = append(res, []int{nums[first], nums[left], nums[right], nums[last]})
				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}

		}
		if maxVal <= target {
			first++
		} else {
			last--
		}
		//first++
		//last--
	}

	m := make(map[string][]int, 0)
	for i := range res {
		slices.Sort(res[i])
		str := ""
		for _, val := range res[i] {
			str += string(val)
		}
		m[str] = res[i]
	}

	uniqueRes := make([][]int, 0)
	for _, v := range m {
		uniqueRes = append(uniqueRes, v)
	}

	return uniqueRes
}

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func min(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

func main() {
	nums := []int{-3, -2, -1, 0, 0, 1, 2, 3}
	target := 0

	res := fourSum(nums, target)

	for i := range res {
		for j := range res[i] {
			print(res[i][j])
			print(" ")
		}
		println()
	}
}
