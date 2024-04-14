package main

import (
	"fmt"
	"strconv"
	"strings"
)

func dayOfYear(date string) int {
	var res int

	nums := strings.Split(date, "-")
	year := nums[0]
	month := nums[1]
	day := nums[2]

	y, _ := strconv.Atoi(year)
	m, _ := strconv.Atoi(month)
	d, _ := strconv.Atoi(day)

	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if y%400 == 0 || (y%4 == 0 && y%100 != 0) {
		days[1] += 1
	}

	for i := 1; i < m; i++ {
		res += days[i-1]
	}

	res += int(d)

	return res
}

func main() {
	date := "2019-01-10"
	res := dayOfYear(date)
	fmt.Println(res)
}
