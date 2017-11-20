package training4_5

func filter(strings []string) []string {
	n := len(strings)
	i := 1
	for i < n {
		if strings[i-1] == strings[i] {
			copy(strings[i-1:], strings[i:])
			n--
		} else {
			i++
		}
	}
	return strings[:n]
}