package main

import "fmt"

func checkOnesSegment(s string) bool {
	zero := false
	for _, c := range s {
		if c == '1' && zero == true {
			return false
		}
		if c == '0' {
			zero = true
		}
	}

	return true
}

func main() {
	fmt.Println(checkOnesSegment("1001"))
	fmt.Println(checkOnesSegment("110"))
}
