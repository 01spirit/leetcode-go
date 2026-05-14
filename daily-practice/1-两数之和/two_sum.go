package main

func twoSum(nums []int, target int) []int {
	sum := make(map[int]int, 0)
	//sum[target-nums[0]] = 0
	//
	//for i := 1; i < len(nums); i++ {
	//	if sum[nums[i]] > 0 || nums[0]+nums[i] == target {
	//		return []int{sum[nums[i]], i}
	//	} else {
	//		sum[target-nums[i]] = i
	//	}
	//}

	for i, num := range nums {
		if p, ok := sum[nums[i]]; ok {
			return []int{p, i}
		}
		sum[target-num] = i
	}

	return nil
}

func main() {
	nums := make([]int, 0)
	//nums = []int{2, 7, 11, 15}
	nums = []int{3, 3}
	target := 6

	res := twoSum(nums, target)
	println(res[0], res[1])
}
