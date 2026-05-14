package main

var ref map[string]string = map[string]string{ // 生命并初始化全局变量
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var results []string // 声明全局变量，用之前需要先初始化

func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return nil
	}

	results = []string{} // 初始化全局变量
	str := ""
	backTrace(digits, 0, str)

	return results
}

/*
	回溯算法
	@param digits : 完整的输入数字字符串
	@param resArr : 结果数组的指针
	@param index : 当前所需处理的 digits 中的数字字符的索引
	@param str : 表示一个组合结果的字符串

	依次向结果字符串中加入字母，再递归调用该函数，处理下一个数字字符的字母，当索引指向输入字符串的末尾时完成一次组合，存储结果，回退到上一个数字字符的处理
	删掉之前加入的该数字字符对应的字母，存入下一个字母，继续递归调用，直到完成所有组合
*/

func backTrace(digits string, index int, str string) {
	if index == len(digits) { // 对字符串完成了一次组合
		//println(str)
		results = append(results, str) // 存入结果数组
	} else {
		ch := digits[index : index+1] // 当前要处理的数字字符
		for i := range ref[ch] {      // 遍历该字符对应的 map 的元素（字母）
			str += string(ref[ch][i])       // 向结果字符串中加入当前数字字符对应的一个字母
			backTrace(digits, index+1, str) // 继续处理下一个数字字符
			str = str[:len(str)-1]          // 完成了一个字母的所有组合，从结果字符串中移除该字母，进入下一个循环，加入下一个字母
		}
	}
}

func main() {
	digits := "23"

	res := letterCombinations(digits)

	for _, str := range res {
		println(str)
	}
}
