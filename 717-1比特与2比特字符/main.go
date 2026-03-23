package main

import "fmt"

func isOneBitCharacter(bits []int) bool {
	one := false
	zero := false

	for i := 0; i < len(bits); i++ {
		if bits[i] == 1 {
			if one {
				one = false
			} else {
				one = true
			}
		} else {
			if one {
				one = false
				zero = false
			} else {
				zero = true
			}
		}
	}

	return zero
}

func main() {
	fmt.Println(isOneBitCharacter([]int{1, 0, 0}))
	fmt.Println(isOneBitCharacter([]int{1, 1, 1, 0}))
}
