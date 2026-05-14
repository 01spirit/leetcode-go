package main

func convert(s string, numRows int) string {
	var res string
	matrix := make([][]int32, numRows)

	if numRows == 1 {
		return s
	}
	length := len(s)
	tol := length*2 - 2
	
	for i := range matrix {
		matrix[i] = make([]int32, tol)
	}

	row := 0
	col := 0
	index := 0
	flag := false
	for _, ch := range s {
		if !flag {
			matrix[row][col] = ch
			row++
			index++
			if index == numRows {
				flag = true
				col++
				row -= 2
				index -= 2
			}
		} else {
			matrix[row][col] = ch
			col++
			row--
			index--
			if index == -1 {
				flag = false
				row += 2
				index += 2
			}
		}
	}

	for i := range matrix {
		for j := range matrix[i] {
			//print(string(matrix[i][j]))
			if matrix[i][j] != '_' {
				res += string(matrix[i][j])
			}
		}
		//println()
	}

	return res
}

func main() {
	s := "Pasfgasgb"
	numRow := 4

	res := convert(s, numRow)

	println(res)

}
