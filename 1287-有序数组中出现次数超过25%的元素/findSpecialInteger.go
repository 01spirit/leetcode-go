package main

import "fmt"

func findSpecialInteger(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	cnt := 1
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {
			cnt++
			if cnt > len(arr)/4 {
				return arr[i]
			}
		} else {
			cnt = 1
		}
	}

	return arr[len(arr)-1]
}

func main() {
	fmt.Println(findSpecialInteger([]int{1, 2, 3, 3}))
	fmt.Println(findSpecialInteger([]int{1, 2, 2, 6, 6, 6, 6, 7, 10}))
}
