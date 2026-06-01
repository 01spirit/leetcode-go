package main

import (
	"fmt"
	"sort"
)

func asteroidsDestroyed(mass int, asteroids []int) bool {
	sort.Ints(asteroids)
	for i := 0; i < len(asteroids); i++ {
		if mass < asteroids[i] {
			return false
		}
		mass += asteroids[i]
	}
	return true
}

func main() {
	fmt.Println(asteroidsDestroyed(10, []int{3, 9, 19, 5, 21}))
	fmt.Println(asteroidsDestroyed(5, []int{4, 9, 23, 4}))
}
