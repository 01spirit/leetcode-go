package main

import "fmt"

func differenceOfSum(nums []int) int {
	sum := 0
	bitSum := 0

	getBitSum := func(num int) (bitSum int) {
		for num > 0 {
			bitSum += num % 10
			num /= 10
		}
		return bitSum
	}

	for _, num := range nums {
		sum += num
		bitSum += getBitSum(num)
	}

	abs := func(n1, n2 int) int {
		if n1 > n2 {
			return n1 - n2
		}
		return n2 - n1
	}

	return abs(sum, bitSum)
}

func main() {
	fmt.Println(differenceOfSum([]int{1, 15, 6, 3}))
	fmt.Println(differenceOfSum([]int{1, 2, 3, 4}))
}
