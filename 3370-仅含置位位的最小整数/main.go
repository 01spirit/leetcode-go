package main

import (
	"fmt"
	"math/bits"
)

func smallestNumber(n int) int {
	ans := 1

	b := bits.Len(uint(n)) - 1
	for i := 0; i < b; i++ {
		ans = ans<<1 | 1
	}

	return ans
}

func main() {
	fmt.Println(smallestNumber(5))
	fmt.Println(smallestNumber(10))
	fmt.Println(smallestNumber(3))
}
