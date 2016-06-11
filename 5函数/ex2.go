package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin) // may add implicit elements such as html, head, or body
	if err != nil {
		log.Fatal(err)
	}
	var counts = make(map[string]int)
	count(counts, doc)
	fmt.Println(counts)
}

// count populates a mapping of element names to the number of elements with that name.
func count(m map[string]int, n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		m[n.Data]++
	}
	count(m, n.FirstChild)
	count(m, n.NextSibling)
}
