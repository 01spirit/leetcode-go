package main

import "fmt"

func getSneakyNumbers(nums []int) []int {
	ex := make([]bool, len(nums))
	ans := make([]int, 0)
	for _, num := range nums {
		if !ex[num] {
			ex[num] = true
		} else {
			ans = append(ans, num)
		}
	}

	return ans
}

func main() {
	fmt.Println(getSneakyNumbers([]int{0, 1, 1, 0}))
	fmt.Println(getSneakyNumbers([]int{0, 3, 2, 1, 3, 2}))
	fmt.Println(getSneakyNumbers([]int{7, 1, 5, 4, 3, 4, 6, 0, 9, 5, 8, 2}))
}
