package training5_15

func min1(vals ...int) int {
	var min int
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}

func min2(val int, vals ...int) int {
	min := val
	for _, v := range vals {
		if v < min {
			min = v
		}
	}
	return min
}