package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int) // counts of Unicode characters
	var utflen [utf8.UTFMax]int  // count of lengths of UTF-8 encodings
	invalid := 0                 // count of invalid UTF-8 characters
	cats := make(map[string]int) // counts of Unicode categories

	// In a terminal, use CTRL+Z at line start to signal EOF with ENTER.
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		switch {
		case unicode.IsLetter(r):
			cats["letter"]++
		case unicode.IsDigit(r):
			cats["digit"]++
		case unicode.IsControl(r):
			cats["control"]++
		case unicode.IsMark(r):
			cats["mark"]++
		case unicode.IsPunct(r):
			cats["punct"]++
		case unicode.IsSymbol(r):
			cats["symbol"]++
		}
		counts[r]++
		utflen[n-1]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		fmt.Printf("%d\t%d\n", i+1, n)
	}
	fmt.Print("\ncat\tcount\n")
	for s, n := range cats {
		fmt.Printf("%v\t%d\n", s, n)
	}
	fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
}
