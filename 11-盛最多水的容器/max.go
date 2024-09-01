package main

/*
设置首尾双指针
*/
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1

	if len(height) == 2 {
		return min(height[0], height[1])
	}

	res := (right - left) * min(height[left], height[right])
	for left < right {
		cur := (right - left) * min(height[left], height[right])
		if res < cur {
			res = cur
		}
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

func min(x, y int) int {
	if x >= y {
		return y
	} else {
		return x
	}
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	res := maxArea(height)

	println(res)
}
