package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	best := math.MaxInt32

	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		left, right := first+1, n-1
		for left < right {
			sum := nums[first] + nums[left] + nums[right]
			if sum == target {
				return target
			}

			best = update(sum, best, target)

			if sum > target {
				tmpIndex := right - 1
				for left < tmpIndex && nums[tmpIndex] == nums[right] {
					tmpIndex--
				}
				right = tmpIndex
			} else {
				tmpIndex := left + 1
				for tmpIndex < right && nums[tmpIndex] == nums[left] {
					tmpIndex++
				}
				left = tmpIndex
			}
		}
	}

	return best
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func update(cur int, best int, target int) int {
	if abs(cur-target) < abs(best-target) {
		return cur
	}
	return best
}

func main() {
	nums := []int{4, 0, 5, -5, 3, 3, 0, -4, -5}
	target := -2

	res := threeSumClosest(nums, target)

	println(res)
}
