package main

func findChampion(grid [][]int) int {
	cmp := -1

	//for i := range grid {
	//	for j := range grid[i] {
	//		if i != j && grid[i][j] == 1 {
	//			if cmp == -1 {
	//				cmp = i
	//			} else {
	//				if cmp == j {
	//					cmp = i
	//				}
	//			}
	//		}
	//	}
	//}

	for i := range grid {
		cnt := 0
		for j := range grid[i] {
			if i != j && grid[i][j] == 0 {
				break
			}
			if i != j && grid[i][j] == 1 {
				cnt++
			}
			if cnt == len(grid[0])-1 {
				return i
			}
		}
	}

	return cmp
}

func main() {
	grid := [][]int{
		{0, 0, 0},
		{1, 0, 0},
		{1, 1, 0},
	}

	res := findChampion(grid)

	println(res)
}
