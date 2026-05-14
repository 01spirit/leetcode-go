package main

import (
	"fmt"
)

func getPrimBitsNum(num int) int {
	cnt := 0

	for num > 0 {
		cnt += num & 1
		num >>= 1
	}

	return cnt
}

func isPrim(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func countPrimeSetBits(left int, right int) int {
	ans := 0
	for i := left; i <= right; i++ {
		cnt := getPrimBitsNum(i)
		if isPrim(cnt) {
			ans++
		}
	}

	return ans
}

func main() {
	fmt.Println(countPrimeSetBits(6, 10))
	fmt.Println(countPrimeSetBits(10, 15))
}
