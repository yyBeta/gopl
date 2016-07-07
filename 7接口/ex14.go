package main

import (
	"fmt"

	"gopl/7接口/eval"
)

func main() {
	expr, _ := eval.Parse("pow(x,min(1.5,2))")
	fmt.Println(expr.Eval(eval.Env{"x": 4}))
}
