package main

import (
	"flag"
	"fmt"
	"gopl/7接口/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
	// the same as 3tempflag.go
	// Kelvin problem done in package tempconv
	// just run with arg such as"-temp 373.15K"
}
