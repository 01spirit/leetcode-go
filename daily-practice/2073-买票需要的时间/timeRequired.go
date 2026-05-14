package main

import "fmt"

func timeRequiredToBuy(tickets []int, k int) int {
	ans := 0

	//for tickets[k] != 0 {
	//	for i := range tickets {
	//		if tickets[i] != 0 {
	//			tickets[i]--
	//			ans++
	//		}
	//		if i == k && tickets[i] == 0 {
	//			break
	//		}
	//	}
	//}

	for i := range tickets {
		if i <= k {
			ans += min(tickets[i], tickets[k])
		} else if i > k {
			ans += min(tickets[i], tickets[k]-1)
		}
	}

	return ans
}

func main() {
	fmt.Println(timeRequiredToBuy([]int{2, 3, 2}, 2))
	fmt.Println(timeRequiredToBuy([]int{5, 1, 1, 1}, 0))
}
