package main

import (
	"fmt"
	"gopl/4复合数据类型/github"
	"log"
	"os"
	"time"
)

const (
	Reset   = "\x1b[0m"
	Black   = "\x1b[30;1m"
	Red     = "\x1b[31;1m"
	Green   = "\x1b[32;1m"
	Yellow  = "\x1b[33;1m"
	Blue    = "\x1b[34;1m"
	Magenta = "\x1b[35;1m"
	Cyan    = "\x1b[36;1m"
	White   = "\x1b[37;1m"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	aYearAgo := time.Now().UTC().AddDate(-1, 0, 0)
	aMonthAgo := time.Now().UTC().AddDate(0, -1, 0)

	for _, item := range result.Items {
		switch {
		case item.CreatedAt.Before(aYearAgo):
			fmt.Print(White)
		case item.CreatedAt.Before(aMonthAgo):
			fmt.Print(Cyan)
		default:
			fmt.Print(Blue)
		}
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
		fmt.Println(Reset)
	}
}
