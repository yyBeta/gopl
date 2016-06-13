package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprint(os.Stderr, "usage: ./ex8.go http://example.com ID_FOR_SEARCH")
		os.Exit(1)
	}
	if err := realmain(os.Args[1], os.Args[2]); err != nil {
		fmt.Fprintf(os.Stderr, "err: %s", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func realmain(url, id string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}
	node := ElementByID(doc, id)
	fmt.Printf("node: %#v", node)
	return nil
}

func ElementByID(doc *html.Node, id string) *html.Node {
	forEachNode(doc, id, startElement, endElement)
	return doc
}

func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) {
	if pre(n, id) {
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, id, pre, post)
	}

	if post(n, id) {
		return
	}
}

var depth int

func startElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				fmt.Print("true!\n")
				return true
			}
		}
	}
	return false
}

func endElement(n *html.Node, id string) bool {
	return n.FirstChild == nil
}
