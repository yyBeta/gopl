package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "textVal: %v\n", err)
		os.Exit(1)
	}
	text(doc)
}

func text(n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.TextNode && n.Data != "script" && n.Data != "style" {
		fmt.Printf("%s\n", n.Attribute)
	}
	text(n.FirstChild)
	text(n.NextSibling)
}
