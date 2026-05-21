package main

func sortColors(nums []int) {
	p0, p1 := 0, 0
	for i, x := range nums {
		nums[i] = 2
		if x <= 1 {
			nums[p1] = 1
			p1++
		}
		if x == 0 {
			nums[p0] = 0
			p0++
		}
	}
}

func main() {
	//fmt.Println(sortColors([]int{2, 0, 2, 1, 1, 0}))
	//fmt.Println(sortColors([]int{2, 0, 1}))
}
