package main

import "fmt"

func temperatureTrend(temperatureA []int, temperatureB []int) int {
	mc := 0

	a := make([]int, len(temperatureA))
	b := make([]int, len(temperatureB))
	cnt := 0
	for i := 0; i < len(temperatureA)-1; i++ {
		a[i] = trend(temperatureA[i+1], temperatureA[i])
		b[i] = trend(temperatureB[i+1], temperatureB[i])
		if a[i] == b[i] {
			cnt++
		} else {
			mc = max(mc, cnt)
			cnt = 0
		}
	}
	mc = max(mc, cnt)

	return mc
}

func trend(a int, b int) int {
	n := a - b
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	} else {
		return 2
	}
}

func main() {
	fmt.Println(temperatureTrend([]int{21, 18, 18, 18, 31}, []int{34, 32, 16, 16, 17}))
	fmt.Println(temperatureTrend([]int{5, 10, 16, -6, 15, 11, 3}, []int{16, 22, 23, 23, 25, 3, -16}))
}
