package main

import "fmt"

func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}
	var ans = 0

	cnt := 0
	for n > 0 {
		bit := (n & 1) ^ 1
		n >>= 1

		ans += bit << cnt
		cnt++
	}

	return ans
}

func main() {
	fmt.Println(bitwiseComplement(5))
	fmt.Println(bitwiseComplement(7))
	fmt.Println(bitwiseComplement(10))
}
