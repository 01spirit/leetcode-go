package main

import "fmt"

func squareIsWhite(coordinates string) bool {
	n1 := coordinates[0] - 'a'
	n2 := coordinates[1] - '1'
	if n1%2 == 0 && n2%2 == 0 || n1%2 == 1 && n2%2 == 1 {
		return false
	}
	return true
}

func main() {
	fmt.Println(squareIsWhite("b2"))
	fmt.Println(squareIsWhite("a1"))
	fmt.Println(squareIsWhite("h3"))
	fmt.Println(squareIsWhite("c7"))
}
