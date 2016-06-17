package main

import "fmt"
import "golang.org/x/net/html"
import "net/http"
import "os"

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var nodes []*html.Node

	if doc.Type == html.ElementNode {
		for _, n := range name {
			if n == doc.Data {
				nodes = append(nodes, doc)
			}
		}
	}

	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, ElementsByTagName(c, name...)...)
	}

	return nodes
}

func fetch(rawurl string) (*html.Node, error) {
	resp, err := http.Get(rawurl)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return html.Parse(resp.Body)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprint(os.Stderr, "usage: ./ex17.go http://example.com h1 h2 h3\n")
		os.Exit(1)
	}
	node, err := fetch(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
	nodes := ElementsByTagName(node, os.Args[2:]...)
	for _, n := range nodes {
		fmt.Printf("%#v\n", n)
	}
}
