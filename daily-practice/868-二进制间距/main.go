package main

import "fmt"

func binaryGap(n int) int {
	ans := 0

	cnt := 0
	begin := false
	for n > 0 {
		bit := n & 1
		n >>= 1
		cnt++
		if bit == 1 {
			if !begin {
				begin = true
			} else {
				ans = max(ans, cnt)
			}
			cnt = 0
		}
	}

	return ans
}

func main() {
	fmt.Println(binaryGap(22))
	fmt.Println(binaryGap(8))
	fmt.Println(binaryGap(5))
}
