package main

import (
	"fmt"
	"gopl/4复合数据类型/github"
	"log"
	"os"
	"time"
)

func main() {
	// 查询issues
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	// 用map存储分好类的issues
	itemByTime := make(map[string][]*github.Issue)
	for _, item := range result.Items {
		if time.Since(item.CreatedAt).Hours() < 24*30 {
			itemByTime["in a month"] = append(itemByTime["in a month"], item)
		} else if time.Since(item.CreatedAt).Hours() < 365*30 {
			itemByTime["from a month to a year"] = append(itemByTime["from a month to a year"], item)
		} else {
			itemByTime["over a year"] = append(itemByTime["over a year"], item)
		}
	}

	// 遍历map输出issues
	for issueType, items := range itemByTime {
		fmt.Println("**** Issues", issueType, "****")
		for _, item := range items {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
}
