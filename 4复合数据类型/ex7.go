package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(str []byte) {
	for i, j, s1, s2, len1, len2 := 0, len(str), rune(0), rune(0), 0, 0; i < j-1; i, j = i+len2, j-len1 {
		s1, len1 = utf8.DecodeRune(str[i:])
		s2, len2 = utf8.DecodeLastRune(str[:j])
		copy(str[i+len2:j-len1], str[i+len1:j-len2])
		copy(str[j-len1:j], []byte(string(s1)))
		copy(str[i:i+len2], []byte(string(s2)))
	}
}

func main() {
	s1 := []byte("Hello, 世界")
	reverse(s1)
	fmt.Println(string(s1))
	s2 := []byte("世界，Hello!")
	reverse(s2)
	fmt.Println(string(s2))
}
