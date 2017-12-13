package training7_2

import "io"

type counter struct {
	writer io.Writer
	writeByte int64
}

func (c *counter) Write(p []byte) (int, error) {
	i, err := c.writer.Write(p)
	if err != nil {
		return 0, err
	}
	c.writeByte += int64(i)

	return len(p), nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := counter{w, 0}
	return &c, &c.writeByte
}