package main

import (
	"fmt"
	"math/bits"
)

func reverseBits(n int) int {
	return int(bits.Reverse32(uint32(n)))
}

func main() {
	fmt.Println(reverseBits(43261596))
	fmt.Println(reverseBits(2147483644))
}
