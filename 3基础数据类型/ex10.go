package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("520"))
	fmt.Println(comma("520184"))
	fmt.Println(comma("5201314"))
}

func comma(s string) string {
	var buf bytes.Buffer
	if len(s) <= 3 {
		return s
	}
	i := len(s) % 3
	buf.WriteString(s[:i])
	for ; i < len(s); i += 3 {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}
