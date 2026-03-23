package main

import "fmt"

func countOdds(low int, high int) int {
	if low%2 == 0 && high%2 == 0 {
		return (high - low) / 2
	}
	return (high-low)/2 + 1
}

func main() {
	fmt.Println(countOdds(3, 7))
	fmt.Println(countOdds(8, 10))
}
