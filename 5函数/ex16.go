package main

import "fmt"

func join(sep string, str ...string) string {
	var strOut string
	length := len(str)
	for i, s := range str {
		if i < length-1 {
			strOut += s + sep
		} else {
			strOut += s
		}
	}
	return strOut
}

func main() {
	fmt.Println(join("/", "hello", "world"))
}
