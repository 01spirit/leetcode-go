package main

func moveZeroes(nums []int) {
	n := len(nums)
	left, right := 0, 0
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

func main() {
	//fmt.Println(moveZeroes([]int{0, 1, 0, 3, 12}))
	//fmt.Println(moveZeroes([]int{0}))
}
