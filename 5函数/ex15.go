package main

import "fmt"

func max(vals ...int) int {
	var m int
	for i, val := range vals {
		if i == 0 {
			m = val
		} else if val > m {
			m = val
		}
	}
	return m
}

func min(vals ...int) int {
	var m int
	for i, val := range vals {
		if i == 0 {
			m = val
		} else if val < m {
			m = val
		}
	}
	return m
}

func main() {
	fmt.Println(max(1, 3, 2))
	fmt.Println(max())
}
