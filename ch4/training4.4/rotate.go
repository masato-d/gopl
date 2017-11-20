package training4_4

func rotate(s []int, i int) []int {
	rtn := s[i:]
	rtn = append(rtn, s[:i]...)
	return rtn
}