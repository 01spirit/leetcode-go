package main

func getCommon(nums1 []int, nums2 []int) int {
	p1, p2 := 0, 0
	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] < nums2[p2] {
			p1++
		} else if nums1[p1] > nums2[p2] {
			p2++
		} else {
			return nums1[p1]
		}
	}
	for p1 < len(nums1) {
		if nums1[p1] == nums2[p2-1] {
			return nums1[p1]
		}

		p1++
	}
	for p2 < len(nums2) {
		if nums2[p2] == nums1[p1-1] {
			return nums2[p2]
		}

		p2++
	}
	return -1
}
