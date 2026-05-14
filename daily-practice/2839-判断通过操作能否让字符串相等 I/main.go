package main

func canBeEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	s10 := swap(s1, 0, 2)
	s11 := swap(s1, 1, 3)
	s12 := swap(s10, 1, 3)

	return s10 == s2 || s11 == s2 || s12 == s2
}

func swap(s string, i int, j int) string {
	b := []byte(s)
	b[i], b[j] = b[j], b[i]
	return string(b)
}

func main() {

}
