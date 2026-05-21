package main

import (
	"fmt"
	"strconv"
)

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	m := map[int]bool{}
	for _, num := range arr1 {
		for num > 0 {
			m[num] = true
			num /= 10
		}
	}
	maxLen := 0
	for _, num := range arr2 {
		for num > 0 {
			if m[num] == true {
				maxLen = max(maxLen, num)
				break
			}
			num /= 10
		}
	}
	if maxLen == 0 {
		return 0
	}
	return len(strconv.Itoa(maxLen))
}

func main() {
	fmt.Println(longestCommonPrefix([]int{1, 10, 100}, []int{1000}))
	fmt.Println(longestCommonPrefix([]int{1, 2, 3}, []int{4, 4, 4}))
}
