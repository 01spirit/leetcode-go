package main

import "fmt"

func removeStars(s string) string {
	result := make([]int32, 0)

	for _, ch := range s {
		if ch == '*' {
			result = result[:len(result)-1]
		} else {
			result = append(result, ch)
		}
	}

	return string(result)
}

func main() {
	fmt.Println(removeStars("leet**cod*e"))
	fmt.Println(removeStars("erase*****"))
}
