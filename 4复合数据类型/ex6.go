package main

import (
	"fmt"
	"unicode"
)

func oneSpace(str []byte) []byte {
	var temp byte
	i := 0
	for _, s := range str {
		if !unicode.IsSpace(rune(s)) || s != temp {
			str[i] = s
			i++
		}
		temp = s
	}
	return str[:i]
}

func main() {
	x := []byte(" hey what's up   hello  world !")
	fmt.Printf("%s", oneSpace(x))
}
