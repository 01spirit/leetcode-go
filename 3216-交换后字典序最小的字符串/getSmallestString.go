package main

import "fmt"

func getSmallestString(s string) string {
	bts := []byte(s)
	for i := range bts {
		if i == len(bts)-1 {
			break
		}
		if (bts[i]-'0')%2 == (bts[i+1]-'0')%2 && bts[i] > bts[i+1] {
			bts[i], bts[i+1] = bts[i+1], bts[i]
			break
		}
	}
	return string(bts)
}

func main() {
	fmt.Println(getSmallestString("45320"))
	fmt.Println(getSmallestString("001"))
}
