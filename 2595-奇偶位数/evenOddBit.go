package main

import "fmt"

func evenOddBit(n int) []int {
	res := []int{0, 0}

	i := 0
	for n > 0 {
		res[i] += n & 1
		n >>= 1
		i ^= 1
	}

	return res
}

func main() {
	fmt.Println(evenOddBit(50))
	fmt.Println(evenOddBit(2))
}
