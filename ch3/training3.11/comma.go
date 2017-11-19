package training3_11

import (
	"strings"
)

func comma(s string) string {
	sl := strings.Split(s, ".")
	sp := ""
	if len(sl) > 1 {
		s = sl[0]
		sp = "." + sl[1]
	}

	n := len(s)
	if n <= 3 {
		return s + sp
	}
	return comma(s[:n-3]) + "," + s[n-3:] + sp
}