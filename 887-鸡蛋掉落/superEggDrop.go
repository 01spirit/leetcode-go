package main

import (
	"fmt"
)

var memo map[int]int = *new(map[int]int)

func superEggDrop(k int, n int) int {
	memo = make(map[int]int)
	return dp(k, n)
}

func dp(k, n int) int {
	if _, ok := memo[n*100+k]; !ok {
		ans := 0
		if n == 0 {
			ans = 0
		} else if k == 1 {
			ans = n
		} else {
			low, high := 1, n
			for low+1 < high {
				mid := (low + high) / 2
				t1 := dp(k-1, mid-1)
				t2 := dp(k, n-mid)
				if t1 < t2 {
					low = mid
				} else if t1 > t2 {
					high = mid
				} else {
					low = mid
					high = mid
				}
			}

			ans = 1 + min(max(dp(k-1, low-1), dp(k, n-low)), max(dp(k-1, high-1), dp(k, n-high)))
		}

		memo[n*100+k] = ans
	}
	return memo[n*100+k]
}

func main() {
	fmt.Println(superEggDrop(3, 8))
	fmt.Println(superEggDrop(1, 2))
	fmt.Println(superEggDrop(2, 6))
	fmt.Println(superEggDrop(3, 14))
}
