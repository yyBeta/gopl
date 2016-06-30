package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int

type LineCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		*c++
	}
	return len(p), nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		*c++
	}
	return len(p), nil
}

func main() {
	var wc WordCounter
	wc.Write([]byte("Hello world! 你好 世界"))
	fmt.Println(wc)
	var lc LineCounter
	lc.Write([]byte(`hello
world`))
	fmt.Println(lc)
}
