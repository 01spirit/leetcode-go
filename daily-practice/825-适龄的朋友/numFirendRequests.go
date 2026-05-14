package main

import (
	"fmt"
)

func numFriendRequests(ages []int) int {
	cnt := 0

	nums := make([]int, 121)
	for _, age := range ages {
		nums[age]++
	}
	nnum := make([][2]int, 0)
	hundred := -1
	for i, num := range nums {
		if num > 0 {
			nnum = append(nnum, [2]int{i, num})
		}
		if hundred == -1 && i >= 100 && num > 0 {
			hundred = len(nnum) - 1
		}
	}

	for i, cur := range nnum {
		if cur[0] > cur[0]/2+7 {
			cnt += cur[1] * (cur[1] - 1)
			//for i := 0; i < cur[1]*(cur[1]-1); i++ {
			//	fmt.Printf("%d - %d\n", cur[0], cur[0])
			//}
		}
		for j := 0; j < i; j++ {
			if nnum[j][0] > nnum[i][0]/2+7 {
				cnt += cur[1] * nnum[j][1]
				fmt.Printf("%d - %d\n", cur[0], nnum[j][0])
			}
		}
	}

	//sort.Ints(ages)
	//for i := 0; i < len(ages); i++ {
	//	for j := 0; j < len(ages); j++ {
	//		if i == j || (ages[j] <= ages[i]/2+7 || ages[j] > ages[i] || (ages[j] > 100 && ages[i] < 100)) {
	//			continue
	//		} else {
	//			cnt++
	//			fmt.Printf("%d - %d\n", ages[i], ages[j])
	//		}
	//	}
	//}

	return cnt
}

func main() {
	fmt.Println(numFriendRequests([]int{73, 106, 39, 6, 26, 15, 30, 100, 71, 35, 46, 112, 6, 60, 110}))
	fmt.Println(numFriendRequests([]int{101, 56, 69, 48, 30}))
	fmt.Println(numFriendRequests([]int{16, 16}))
	fmt.Println(numFriendRequests([]int{16, 17, 18}))
	fmt.Println(numFriendRequests([]int{20, 30, 100, 110, 120}))
}
