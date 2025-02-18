package main

type RangeFreqQuery struct {
	occ map[int][]int
}

func Constructor(arr []int) RangeFreqQuery {
	occ := make(map[int][]int)
	for k, v := range arr {
		occ[v] = append(occ[v], k)
	}
	return RangeFreqQuery{occ}
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	pos, exit := this.occ[value]
	if !exit {
		return 0
	}
	l := lowerBound(pos, left)
	r := upperBound(pos, right)
	return r - l
}

func lowerBound(pos []int, target int) int {
	low, high := 0, len(pos)-1
	for low <= high {
		mid := low + (high-low)>>1
		if pos[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

func upperBound(pos []int, target int) int {
	low, high := 0, len(pos)-1
	for low <= high {
		mid := low + (high-low)>>1
		if pos[mid] <= target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */
