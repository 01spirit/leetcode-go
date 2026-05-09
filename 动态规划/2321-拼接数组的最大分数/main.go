package main

import "fmt"

func maxSubArray(nums1, nums2 []int) int {
	ans := 0
	preSum := 0
	maxSum := 0

	for i := 0; i < len(nums1); i++ {
		preSum += nums1[i]
		ans = max(ans, 0) + nums2[i] - nums1[i]
		maxSum = max(maxSum, ans)
	}

	return preSum + maxSum
}

func maximumsSplicedArray(nums1 []int, nums2 []int) int {
	return max(maxSubArray(nums1, nums2), maxSubArray(nums2, nums1))
}

func main() {
	fmt.Println(maximumsSplicedArray([]int{60, 60, 60}, []int{10, 90, 10}))
	fmt.Println(maximumsSplicedArray([]int{20, 40, 20, 70, 30}, []int{50, 20, 50, 40, 20}))
	fmt.Println(maximumsSplicedArray([]int{7, 11, 13}, []int{1, 1, 1}))
}
