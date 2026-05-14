package main

import (
	"fmt"
	"unicode"
)

func clearDigits(s string) string {
	var result []byte
	for _, ch := range s {
		if unicode.IsDigit(ch) {
			result = result[:len(result)-1]
		} else {
			result = append(result, byte(ch))
		}
	}

	return string(result)
}

func main() {
	fmt.Println(clearDigits("cd45"))
	fmt.Println(clearDigits("eraas"))
	fmt.Println(clearDigits("a5s6d7cc89vvv000as"))
}
