package main

import "fmt"

func destCity(paths [][]string) string {
	dest := make(map[string]bool)

	for _, path := range paths {
		dest[path[0]] = true
		if !dest[path[1]] {
			dest[path[1]] = false
		}
	}

	for key := range dest {
		if !dest[key] {
			return key
		}
	}

	return ""
}

func main() {
	fmt.Println(destCity([][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}))
	fmt.Println(destCity([][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}))
	fmt.Println(destCity([][]string{{"A", "Z"}}))
}
