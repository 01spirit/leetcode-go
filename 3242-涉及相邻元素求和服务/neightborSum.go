package main

type NeighborSum struct {
	grid [][]int
}

func Constructor(grid [][]int) NeighborSum {
	return NeighborSum{
		grid: grid,
	}
}

func (this *NeighborSum) AdjacentSum(value int) int {
	res := 0
	for i := 0; i < len(this.grid); i++ {
		for j := 0; j < len(this.grid[i]); j++ {
			if this.grid[i][j] == value {
				if i-1 >= 0 {
					res += this.grid[i-1][j]
				}
				if j-1 >= 0 {
					res += this.grid[i][j-1]
				}
				if i+1 < len(this.grid) {
					res += this.grid[i+1][j]
				}
				if j+1 < len(this.grid[i]) {
					res += this.grid[i][j+1]
				}
				return res
			}
		}
	}

	return res
}

func (this *NeighborSum) DiagonalSum(value int) int {
	res := 0
	for i := 0; i < len(this.grid); i++ {
		for j := 0; j < len(this.grid[i]); j++ {
			if this.grid[i][j] == value {
				if i-1 >= 0 && j-1 >= 0 {
					res += this.grid[i-1][j-1]
				}
				if j-1 >= 0 && i+1 < len(this.grid) {
					res += this.grid[i+1][j-1]
				}
				if i+1 < len(this.grid) && j+1 < len(this.grid[i]) {
					res += this.grid[i+1][j+1]
				}
				if j+1 < len(this.grid[i]) && i-1 >= 0 {
					res += this.grid[i-1][j+1]
				}
				return res
			}
		}
	}

	return res
}
