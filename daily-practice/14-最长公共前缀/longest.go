package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	res := strs[0]
	for i := 1; i < len(strs); i++ {
		index := 0
		if len(strs[i]) == 0 {
			return ""
		}
		length := 0
		if len(res) < len(strs[i]) {
			length = len(res)
		} else {
			length = len(strs[i])
		}
		for index < length {
			if res[index] != strs[i][index] {
				if index == 0 {
					return ""
				} else {
					res = res[:index]
					break
				}
			} else {
				index++
				if index == length {
					res = strs[i][:index]
				}
			}
		}
	}

	return res
}

func main() {
	strs := []string{"ab", "a"}

	res := longestCommonPrefix(strs)

	println(res)
}
