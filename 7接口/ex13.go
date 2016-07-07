package main

import (
	"fmt"

	"gopl/7接口/eval"
)

func main() {
	expr, _ := eval.Parse("pow(x + pow(y))")
	fmt.Println(expr.String())
}
