package main

import "slices"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var res float64

	m := len(nums1)
	n := len(nums2)

	//tmpArr := make([]int, 0)
	nums1 = append(nums1, nums2...)
	slices.Sort(nums1)
	//tlen := len(nums1)

	if (m+n)%2 == 1 {
		res = float64(nums1[(m+n)/2])
		//println(res)
	} else {
		n1 := float64(nums1[(m+n+1)/2])
		n2 := float64(nums1[(m+n-1)/2])
		//println(n1)
		//println(n2)
		res = (n1 + n2) / 2

	}

	return res
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}

	res := findMedianSortedArrays(nums1, nums2)

	println(res)
}
