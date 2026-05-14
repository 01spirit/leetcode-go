package main

func romanToInt(s string) int {
	ref := make(map[string]int, 0)
	ref["I"] = 1
	ref["V"] = 5
	ref["X"] = 10
	ref["L"] = 50
	ref["C"] = 100
	ref["D"] = 500
	ref["M"] = 1000

	ref["IV"] = 4
	ref["IX"] = 9
	ref["XL"] = 40
	ref["XC"] = 90
	ref["CD"] = 400
	ref["CM"] = 900

	res := 0

	for i := 0; i < len(s); i++ {
		if i < len(s)-1 { // 如果不是最后一位，可能出现两个字母的组合
			str := s[i : i+2]
			if val, ok := ref[str]; ok { // 是两个字母的组合
				res += val
				i++
			} else { // 不是字母组合
				res += ref[s[i:i+1]]
			}
		} else { // 最后一位
			res += ref[s[i:i+1]]
		}
	}

	return res
}

func main() {
	s := "MCMXCIV"

	res := romanToInt(s)

	print(res)

}
