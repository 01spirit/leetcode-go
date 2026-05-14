package main

import (
	"fmt"
	"math"
	"slices"
)

func generateKey(num1 int, num2 int, num3 int) int {
	res := 0
	for i := 0; i < 4; i++ {
		n1 := num1 % 10
		n2 := num2 % 10
		n3 := num3 % 10
		num1 /= 10
		num2 /= 10
		num3 /= 10
		//fmt.Println(n1, n2, n3)
		if n1 == 0 || n2 == 0 || n3 == 0 {
			continue
		}
		tmp := []int{n1, n2, n3}
		slices.Sort(tmp)
		res += tmp[0] * int(math.Pow(10, float64(i)))
		//fmt.Println(tmp)

	}
	return res
}

func main() {
	fmt.Println(generateKey(1140, 1851, 2057))
	fmt.Println(generateKey(1, 10, 1000))
	fmt.Println(generateKey(987, 879, 798))
	fmt.Println(generateKey(1, 2, 3))
}
