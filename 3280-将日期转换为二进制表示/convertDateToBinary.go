package main

import (
	"fmt"
	"strconv"
)

func convertDateToBinary(date string) string {
	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:])

	res := ""
	res = fmt.Sprintf("%b-%b-%b", year, month, day)

	return res
}

func main() {
	fmt.Println(convertDateToBinary("2080-02-29"))
	fmt.Println(convertDateToBinary("1900-01-01"))
}
