package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(newComma("+5201314.2333"))
	fmt.Println(newComma("-5201314"))
}

func newComma(s string) string {
	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		return string(s[0]) + newComma(s[1:])
	}
	dot := strings.LastIndex(s, ".")
	if dot == -1 {
		return comma(s)
	}
	return comma(s[:dot]) + s[dot:]
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
