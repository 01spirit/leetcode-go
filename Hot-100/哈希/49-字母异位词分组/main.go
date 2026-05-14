package main

import (
	"fmt"
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	hash := make(map[string][]int)
	tmp := ""
	bt := []byte{}
	for i, str := range strs {
		bt = []byte(str)
		slices.Sort(bt)
		tmp = string(bt)
		if arr, ok := hash[tmp]; ok {
			arr = append(arr, i)
			hash[tmp] = arr
		} else {
			hash[tmp] = []int{}
			hash[tmp] = append(hash[tmp], i)
		}
	}

	ans := make([][]string, 0)
	for _, arr := range hash {
		tmpArr := make([]string, 0)
		for _, i := range arr {
			tmpArr = append(tmpArr, strs[i])
		}
		ans = append(ans, tmpArr)
	}
	return ans
}

//func groupAnagrams(strs []string) [][]string {
//	hash := map[string][]string{}
//	tmp := ""
//	bt := []byte{}
//	for i, str := range strs {
//		bt = []byte(str)
//		slices.Sort(bt)
//		tmp = string(bt)
//		hash[tmp] = append(hash[tmp], strs[i])
//	}
//
//	return slices.Collect(maps.Values(hash))
//}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{""}))
	fmt.Println(groupAnagrams([]string{"a"}))
}
