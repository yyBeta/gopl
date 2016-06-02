package main

import "fmt"

func delDup(str []string) []string {
	temp := ""
	i := 0
	for _, s := range str {
		if s != temp {
			str[i] = s
			i++
		}
		temp = s
	}
	return str[:i]
}

func main() {
	x := []string{"aa", "aa", "bb", "cc", "bb", "bb", "bb", "dd"}
	fmt.Println(delDup(x))
}
