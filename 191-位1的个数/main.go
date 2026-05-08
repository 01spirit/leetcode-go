package main

import "fmt"

func hammingWeight(n int) int {
	ans := 0
	for n > 0 {
		ans += n & 1
		n >>= 1
	}
	return ans
}

func main() {
	fmt.Println(hammingWeight(11))
	fmt.Println(hammingWeight(128))
	fmt.Println(hammingWeight(2147483645))
}
