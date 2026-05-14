package main

import "fmt"

func replaceElements(arr []int) []int {
	answer := make([]int, len(arr))
	answer[len(arr)-1] = -1

	num := arr[len(arr)-1]
	for i := len(arr) - 2; i >= 0; i-- {
		num = max(num, arr[i+1])
		answer[i] = num
	}

	return answer
}

func main() {
	fmt.Println(replaceElements([]int{17, 18, 5, 4, 6, 1}))
	fmt.Println(replaceElements([]int{400}))
}
