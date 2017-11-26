package training5_9

import (
	"strings"
)

func expand(s string, f func(string) string) string {
	ss := strings.Fields(s)
	if len(ss) < 1 {
		return ""
	}

	var r []string
	for _, v := range ss {
		if strings.HasPrefix(v, "$") {
			arg := v[1:]
			if arg != "" {
				r = append(r, f(arg))
			}
		} else {
			r = append(r, v)
		}
	}
	return strings.Join(r, " ")
}