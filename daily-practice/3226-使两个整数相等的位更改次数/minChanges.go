package main

import "fmt"

func minChanges(n int, k int) int {
	if n == k {
		return 0
	}
	ans := 0
	for n > 0 {
		b1 := n & 0x1
		b2 := k & 0x1
		if b1 == 0 && b2 == 1 {
			return -1
		} else if b1 != b2 {
			ans++
		}
		n = n >> 1
		k = k >> 1
	}
	if k != n {
		return -1
	}

	return ans
}

func main() {
	fmt.Println(minChanges(11, 56))
	fmt.Println(minChanges(13, 4))
	fmt.Println(minChanges(21, 21))
	fmt.Println(minChanges(14, 13))
}
