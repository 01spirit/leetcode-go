package main

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	//str := strconv.FormatInt(int64(x), 10)
	//for i := 0; i < len(str)/2; i++ {
	//	if str[i] != str[len(str)-1-i] {
	//		return false
	//	}
	//}

	half := 0
	for half < x {
		half = half*10 + x%10
		x /= 10
	}
	if half == x || half/10 == x {
		return true
	}

	return false
}

func main() {
	x := 101

	res := isPalindrome(x)

	println(res)
}
