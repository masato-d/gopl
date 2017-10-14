package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf1 bytes.Buffer
	var buf2 bytes.Buffer

	n := len(s) - 1
	for i := n; i >= 0; i-- {
		if (n - i) % 3 == 0 && n != i {
			buf1.WriteString(",")
		}
		buf1.WriteByte(s[i])
	}

	for i := buf1.Len()-1; i >= 0; i-- {
		buf2.WriteByte(buf1.Bytes()[i])
	}

	return buf2.String()
}

func main() {
	s := "1234567890"
	fmt.Println(comma(s))
}