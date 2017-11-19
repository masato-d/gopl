package training3_12

import (
	"strings"
)

func anagram(w1, w2 string) bool {
	if len(w1) != len(w2) {
		return false
	}

	checked := w2
	for i := 0; i < len(w1) ; i++ {
		if !strings.Contains(checked, string(w1[i])) {
			return false
		}
		checked = strings.Replace(checked, string(w1[i]), "", 1)
	}

	if len(checked) > 0 {
		return false
	}

	return true
}
