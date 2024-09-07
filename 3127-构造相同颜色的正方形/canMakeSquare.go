package main

import "fmt"

func canMakeSquare(grid [][]byte) bool {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if stats(grid, i, j) {
				return true
			}
		}
	}

	return false
}

func stats(grid [][]byte, i, j int) bool {
	w, b := 0, 0
	if grid[i][j] == 'W' {
		w++
	} else {
		b++
	}
	if grid[i+1][j] == 'W' {
		w++
	} else {
		b++
	}
	if grid[i][j+1] == 'W' {
		w++
	} else {
		b++
	}
	if grid[i+1][j+1] == 'W' {
		w++
	} else {
		b++
	}

	if w >= 3 || b >= 3 {
		return true
	}
	return false
}

func main() {
	fmt.Println(canMakeSquare([][]byte{[]byte{'B', 'W', 'B'}, {'B', 'W', 'W'}, {'B', 'W', 'B'}}))
	fmt.Println(canMakeSquare([][]byte{[]byte{'B', 'W', 'B'}, {'W', 'B', 'W'}, {'B', 'W', 'B'}}))
	fmt.Println(canMakeSquare([][]byte{[]byte{'B', 'W', 'B'}, {'B', 'W', 'W'}, {'B', 'W', 'W'}}))
}
