package main

import (
	"fmt"
	"gopl/6方法/geometry"
)

func main() {
	perim := geometry.Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
	fmt.Println(perim.Distance()) // "12", method of geometry.Path
}
