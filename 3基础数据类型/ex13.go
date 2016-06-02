package main

import "fmt"

const (
	_ = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(EB)
	// 超出int范围
	// fmt.Println(ZB)
	// fmt.Println(YB)
}
