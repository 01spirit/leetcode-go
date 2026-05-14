package main

import (
	"math"
)

//func reverse(x int) int {
//	src := int64(x)
//
//	neg := false
//	if x < 0 {
//		neg = true
//		src = -src
//	}
//	bef := strconv.FormatInt(src, 10)
//	aft := revStr(bef)
//
//	res, err := strconv.ParseInt(aft, 10, 32)
//	//println(aft)
//
//	if err != nil {
//		return 0
//	}
//
//	if neg {
//		return -int(res)
//	}
//
//	return int(res)
//}
//
//func revStr(str string) string {
//	var res string
//	for i := len(str) - 1; i >= 0; i-- {
//		res += string(str[i])
//	}
//	return res
//}

func reverse(x int) int {

	res := 0
	for x != 0 {
		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		res = res*10 + digit
	}

	return res
}

func main() {
	x := -2147483412

	res := reverse(x)

	println(res)

}
