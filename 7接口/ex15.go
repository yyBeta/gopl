package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopl/7接口/eval"
)

const assignment_error = 2

func main() {
	exitCode := 0
	stdin := bufio.NewScanner(os.Stdin)
	fmt.Printf("Expression: ")
	stdin.Scan()
	expr, err := eval.Parse(stdin.Text())
	if err != nil {
		fmt.Fprintf(os.Stderr, "bad expression: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Variables (<var>=<val>, eg: x=3 y=5 ...): ")
	stdin.Scan()
	envStr := stdin.Text()
	if stdin.Err() != nil {
		fmt.Fprintln(os.Stderr, stdin.Err())
		os.Exit(1)
	}

	env := eval.Env{}
	assignments := strings.Fields(envStr)
	for _, a := range assignments {
		fields := strings.Split(a, "=")
		if len(fields) != 2 {
			fmt.Fprintf(os.Stderr, "bad assignment: %s\n", a)
			exitCode = assignment_error
		}
		ident, valStr := fields[0], fields[1]
		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "bad value for %s, using zero: %s\n", ident, err)
			exitCode = assignment_error
		}
		env[eval.Var(ident)] = val
	}

	fmt.Println(expr.Eval(env))
	os.Exit(exitCode)
}
