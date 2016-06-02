package main

import (
	"crypto/sha256"
	"fmt"
)

func diff(x, y [32]byte) int {
	var count int
	for i := 0; i < 32; i++ {
		for j := uint(0); j < 8; j++ {
			if x[i]&byte(1<<j) == y[i]&byte(1<<j) {
				count++
			}
		}
	}
	return count
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%d\n", c1, c2, diff(c1, c2))
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// 131
}
