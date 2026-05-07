package main

import "math"

func furthestDistanceFromOrigin(moves string) int {
	lc, rc := 0, 0
	cc := 0
	for _, m := range moves {
		if m == 'L' {
			lc++
		} else if m == 'R' {
			rc++
		} else {
			cc++
		}
	}
	return int(math.Abs(float64(lc-rc))) + cc
}

func main() {

}
