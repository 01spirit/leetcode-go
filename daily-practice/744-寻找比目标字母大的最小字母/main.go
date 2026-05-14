package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
	if letters[len(letters)-1] <= target {
		return letters[0]
	}
	left := 0
	right := len(letters)

	var ans byte = 0
	for left < right {
		mid := (left + right) / 2
		if letters[mid] <= target {
			left = mid + 1
		} else {
			right = mid
			ans = letters[mid]
		}
	}

	if ans > target {
		return ans
	}

	return letters[0]
}

func main() {
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a')))
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c')))
	fmt.Println(string(nextGreatestLetter([]byte{'x', 'x', 'y', 'y'}, 'z')))
}
