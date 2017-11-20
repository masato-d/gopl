package training4_6

import "unicode"

func compress(str []byte) []byte {
	i := 0
	for _, s := range str {
		if unicode.IsSpace(rune(s)) && unicode.IsSpace(rune(str[i-1])) {
			continue
		}
		str[i] = s
		i++
	}
	return str[:i]
}