// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordFreq := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	//input "ctrl+z" as EOF to end inputing
	for input.Scan() {
		wordFreq[input.Text()]++
	}

	for word, count := range wordFreq {
		fmt.Printf("word:%s\tcount:%d\n", word, count)
	}
}
