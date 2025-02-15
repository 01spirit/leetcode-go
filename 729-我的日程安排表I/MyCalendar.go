package main

import (
	"fmt"
	"sort"
)

type MyCalendar struct {
	workList [][]int
}

func Constructor() MyCalendar {
	return MyCalendar{make([][]int, 0)}
}

func (this *MyCalendar) Book(startTime int, endTime int) bool {
	if len(this.workList) == 0 {
		this.workList = append(this.workList, []int{startTime, endTime})
		return true
	}

	for i := range this.workList {
		if this.workList[i][0] < endTime && this.workList[i][1] > startTime {
			return false
		}
	}

	this.workList = append(this.workList, []int{startTime, endTime})
	sort.Slice(this.workList, func(i, j int) bool {
		return this.workList[i][1] < this.workList[j][1]
	})

	return true
}

func main() {
	mc := Constructor()
	fmt.Println(mc.Book(10, 20))
	fmt.Println(mc.Book(15, 25))
	fmt.Println(mc.Book(20, 30))
}
