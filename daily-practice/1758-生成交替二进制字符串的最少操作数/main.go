package main

import (
	"fmt"
	"math"
)

func minOperations(s string) int {
	ans := 0

	oddOne := 0
	evenOne := 0
	oddZero := 0
	evenZero := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			if i%2 == 0 {
				evenOne++
			} else {
				oddOne++
			}
		} else {
			if i%2 == 0 {
				evenZero++
			} else {
				oddZero++
			}
		}
	}

	ans += int(math.Abs(float64(len(s) - max(oddOne+evenZero, oddZero+evenOne))))

	return ans
}

func main() {
	fmt.Println(minOperations("0100"))
	fmt.Println(minOperations("1111"))
	fmt.Println(minOperations("110010"))
}
