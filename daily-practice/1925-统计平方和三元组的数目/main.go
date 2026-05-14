package main

import (
	"fmt"
	"math"
)

func countTriples(n int) int {
	ans := 0

	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			cc := i*i + j*j
			c := int(math.Sqrt(float64(cc)))
			if cc <= n*n && c*c == cc {
				ans++
			}
		}
	}

	return ans * 2
}

func main() {
	fmt.Println(countTriples(5))
	fmt.Println(countTriples(10))
}
