package main

import "fmt"

func sumDigitDifferences(nums []int) int64 {
	cnt := int64(0)

	for nums[0] > 0 {
		bts := make([]int, 10)
		for i := range nums {
			bts[nums[i]%10]++
			nums[i] /= 10
		}

		for i := 0; i < 10; i++ {
			cnt += int64(bts[i]) * int64(len(nums)-bts[i])
		}
	}

	return cnt / 2
}

func main() {
	fmt.Println(sumDigitDifferences([]int{13, 23, 12}))
	fmt.Println(sumDigitDifferences([]int{10, 10, 10}))
	fmt.Println(sumDigitDifferences([]int{50, 28, 48}))
}
