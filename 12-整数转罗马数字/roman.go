package main

//func intToRoman(num int) string {
//	ref := make(map[int]string, 0)
//	ref[1] = "I"
//	ref[5] = "V"
//	ref[10] = "X"
//	ref[50] = "L"
//	ref[100] = "C"
//	ref[500] = "D"
//	ref[1000] = "M"
//
//	ref[4] = "IV"
//	ref[9] = "IX"
//	ref[40] = "XL"
//	ref[90] = "XC"
//	ref[400] = "CD"
//	ref[900] = "CM"
//
//	res := ""
//	for num > 0 {
//		if num/1000 > 0 {
//			tmp := num / 1000
//			for tmp > 0 {
//				res += ref[1000]
//				tmp--
//			}
//			num %= 1000
//		} else if num/500 > 0 {
//			if num/900 == 1 {
//				res += ref[900]
//				num %= 900
//			} else {
//				tmp := num / 500
//				for tmp > 0 {
//					res += ref[500]
//					tmp--
//				}
//				num %= 500
//			}
//		} else if num/100 > 0 {
//			if num/400 == 1 {
//				res += ref[400]
//				num %= 400
//			} else {
//				tmp := num / 100
//				for tmp > 0 {
//					res += ref[100]
//					tmp--
//				}
//				num %= 100
//			}
//		} else if num/50 > 0 {
//			if num/90 == 1 {
//				res += ref[90]
//				num %= 90
//			} else {
//				tmp := num / 50
//				for tmp > 0 {
//					res += ref[50]
//					tmp--
//				}
//				num %= 50
//			}
//		} else if num/10 > 0 {
//			if num/40 == 1 {
//				res += ref[40]
//				num %= 40
//			} else {
//				tmp := num / 10
//				for tmp > 0 {
//					res += ref[10]
//					tmp--
//				}
//				num %= 10
//			}
//		} else if num/5 > 0 {
//			if num/9 == 1 {
//				res += ref[9]
//				num %= 9
//			} else {
//				tmp := num / 5
//				for tmp > 0 {
//					res += ref[5]
//					tmp--
//				}
//				num %= 5
//			}
//		} else if num > 0 {
//			if num == 4 {
//				res += ref[4]
//				num %= 4
//			} else {
//				tmp := num
//				for tmp > 0 {
//					res += ref[1]
//					tmp--
//					num %= 1
//				}
//			}
//		}
//	}
//
//	return res
//}

var (
	thousands = []string{"", "M", "MM", "MMM"}
	hundreds  = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens      = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones      = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
)

func intToRoman(num int) string {
	return thousands[num/1000] + hundreds[num%1000/100] + tens[num%100/10] + ones[num%10]
}

func main() {
	num := 1994

	res := intToRoman(num)

	println(res)
}
