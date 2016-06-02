package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var bit = flag.Int("bit", 256, "bit of SHA. Default:256.")

func main() {
	flag.Parse()
	var b []byte
	fmt.Printf("input string to see its SHA%d: ", *bit)
	fmt.Scanln(&b)

	switch *bit {
	case 256:
		fmt.Fprintf(os.Stdout, "%x", sha256.Sum256(b))
	case 384:
		fmt.Fprintf(os.Stdout, "%x", sha512.Sum384(b))
	case 512:
		fmt.Fprintf(os.Stdout, "%x", sha512.Sum512(b))
	default:
		fmt.Fprintf(os.Stderr, "unsupported bit")
	}
}

// run as "ex2.exe -bit=384"
