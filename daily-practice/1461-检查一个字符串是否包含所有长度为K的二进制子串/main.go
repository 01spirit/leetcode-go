package main

import (
	"fmt"
	"math"
)

func convert(s string) uint64 {
	b := []byte(s)

	res := uint64(0)
	for i := 0; i < len(b); i++ {
		res = res<<1 | uint64(b[i]-'0')
	}

	return res
}

func hasAllCodes(s string, k int) bool {
	substrings := make(map[uint64]struct{})
	for i := 0; i < len(s)-k+1; i++ {
		substrings[convert(s[i:i+k])] = struct{}{}
	}
	return len(substrings) == int(math.Exp2(float64(k)))
}

func main() {
	fmt.Println(hasAllCodes("00110", 2))
	fmt.Println(hasAllCodes("00110110", 2))
	fmt.Println(hasAllCodes("0110", 1))
	fmt.Println(hasAllCodes("0110", 2))
}
