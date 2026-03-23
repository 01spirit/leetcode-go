package main

import (
	"fmt"
	"slices"
	"unicode"
)

func check(code string, isActive bool) bool {
	if !isActive {
		return false
	}

	for _, c := range code {
		if c != '_' && !unicode.IsLetter(c) && !unicode.IsDigit(c) {
			return false
		}
	}

	return true
}

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	groups := make([][]string, 4)
	for i := range groups {
		groups[i] = make([]string, 0)
	}

	ans := make([]string, 0)

	for i := 0; i < len(code); i++ {
		if code[i] != "" && check(code[i], isActive[i]) {
			switch businessLine[i] {
			case "electronics":
				groups[0] = append(groups[0], code[i])
				break
			case "grocery":
				groups[1] = append(groups[1], code[i])
				break
			case "pharmacy":
				groups[2] = append(groups[2], code[i])
				break
			case "restaurant":
				groups[3] = append(groups[3], code[i])
				break
			}
		}
	}

	for i := 0; i < len(groups); i++ {
		slices.Sort(groups[i])
		ans = append(ans, groups[i]...)
	}

	return ans
}

func main() {
	fmt.Println(validateCoupons([]string{"SAVE20", "", "PHARMA5", "SAVE@20"}, []string{"restaurant", "grocery", "pharmacy", "restaurant"}, []bool{true, true, true, true}))
	fmt.Println(validateCoupons([]string{"GROCERY15", "ELECTRONICS_50", "DISCOUNT10"}, []string{"grocery", "electronics", "invalid"}, []bool{false, true, true}))
}
