package training7_1

import (
	"bufio"
	"bytes"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanWords)

	var count int
	for sc.Scan() {
		count++
	}

	if err := sc.Err(); err != nil {
		return 0, err
	}

	*c = WordCounter(count)

	return len(p), nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewReader(p))

	var count int
	for sc.Scan() {
		count++
	}

	if err := sc.Err(); err != nil {
		return 0, err
	}

	*c = LineCounter(count)

	return len(p), nil
}