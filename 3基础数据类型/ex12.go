package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(strconv.FormatBool(isAnagram("Anagram", "gramanA")))
}

func isAnagram(s1 string, s2 string) bool {
	if s1 == s2 || len(s1) != len(s2) {
		return false
	}
	for _, v := range s1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}
	return true
}
