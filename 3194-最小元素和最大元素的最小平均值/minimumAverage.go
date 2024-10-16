package main

import (
	"fmt"
	"math"
)

// nums 取值 [1, 50] ,可以用计数排序，时间复杂度 O(n)，需要额外开辟两个数组，空间复杂度 O(n+51);
// nums.length <= 50，所以实际时间上并不占优
func minimumAverage(nums []int) float64 {
	cnt := make([]int, 51)

	for _, num := range nums {
		cnt[num]++
	}
	for i := 1; i < len(cnt); i++ {
		cnt[i] += cnt[i-1]
	}
	s := make([]int, len(nums)+1)
	for _, num := range nums {
		s[cnt[num]] = num
		cnt[num]--
	}

	average := float64(math.MaxInt32)
	left, right := 1, len(s)-1
	for left < right {
		average = min(average, float64(s[left]+s[right])/2)
		left++
		right--
	}

	return average
}

func main() {
	fmt.Println(minimumAverage([]int{1, 9, 8, 3, 10, 5}))
	fmt.Println(minimumAverage([]int{7, 8, 3, 4, 15, 13, 4, 1}))
	fmt.Println(minimumAverage([]int{1, 2, 3, 7, 8, 9}))
}
