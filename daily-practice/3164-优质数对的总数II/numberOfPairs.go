package main

import "fmt"

func numberOfPairs(nums1 []int, nums2 []int, k int) int64 {
	ans := int64(0)

	cnt1 := make(map[int]int)
	cnt2 := make(map[int]int)
	hbound := 0
	for _, num := range nums1 {
		cnt1[num]++
		hbound = max(hbound, num)
	}
	for _, num := range nums2 {
		cnt2[num]++
	}

	for num, cnt := range cnt2 {
		for factor := num * k; factor <= hbound; factor += num * k {
			if f, ok := cnt1[factor]; ok {
				ans += int64(f * cnt)
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(numberOfPairs([]int{1, 3, 4}, []int{1, 3, 4}, 1))
	fmt.Println(numberOfPairs([]int{1, 2, 4, 12}, []int{2, 4}, 3))
}
