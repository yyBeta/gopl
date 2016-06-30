package main

import (
	"fmt"
	"io"
	"os"
)

type wrapper struct {
	counter int64
	writer  io.Writer
}

func (c *wrapper) Write(p []byte) (int, error) {
	n, err := c.writer.Write(p)
	c.counter += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	wr := &wrapper{writer: w}
	return wr, &wr.counter
}

func main() {
	writer, count := CountingWriter(os.Stdout)
	fmt.Fprint(writer, "Hello world\n")
	fmt.Println(*count)
}
