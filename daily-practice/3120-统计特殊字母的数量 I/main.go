package main

import "fmt"

func numberOfSpecialChars(word string) int {
	ms := map[byte]bool{}
	mb := map[byte]bool{}
	for _, v := range word {
		if v-'a' >= 0 {
			ms[byte(v-'a')] = true
		} else {
			mb[byte(v-'A')] = true
		}
	}
	cnt := 0
	for v := range mb {
		if _, ok := ms[v]; ok {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(numberOfSpecialChars("aaAbcBC"))
	fmt.Println(numberOfSpecialChars("abc"))
	fmt.Println(numberOfSpecialChars("abBCab"))
}
