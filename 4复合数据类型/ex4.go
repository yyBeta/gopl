package main

import "fmt"

func rotate(ints []int, offset int) {
	temp := make([]int, len(ints))
	for i, s := range ints {
		temp[(len(ints)+i-offset)%len(ints)] = s
	}
	copy(ints, temp)
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	rotate(s, 2)
	fmt.Println(s)
}
