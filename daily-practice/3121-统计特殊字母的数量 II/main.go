package main

import "fmt"

func numberOfSpecialChars(word string) int {
	ms := map[byte]int{}
	mb := map[byte]int{}
	for i, v := range word {
		if v-'a' >= 0 {
			ms[byte(v-'a')] = i
		} else {
			if _, ok := mb[byte(v-'A')]; !ok {
				mb[byte(v-'A')] = i
			}
		}
	}

	cnt := 0
	for v := range mb {
		if s, ok := ms[v]; ok {
			if s < mb[v] {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(numberOfSpecialChars("cCceDC"))
	fmt.Println(numberOfSpecialChars("aaAbcBC"))
	fmt.Println(numberOfSpecialChars("abc"))
	fmt.Println(numberOfSpecialChars("AbBCab"))
}
