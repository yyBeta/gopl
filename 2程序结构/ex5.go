package main

import (
	"fmt"
	"gopl/2程序结构/popcount"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(popcount.PopCountLoop(0x1234567890ABCDEF))
	fmt.Println(time.Since(start).Nanoseconds())

	start = time.Now()
	fmt.Println(popcount.BitCount(0x1234567890ABCDEF))
	fmt.Println(time.Since(start).Nanoseconds())

	start = time.Now()
	fmt.Println(popcount.PopCountByClearing(0x1234567890ABCDEF))
	fmt.Println(time.Since(start).Nanoseconds())
	// 由于还没学testing包，只能尝试用time评价性能
}
